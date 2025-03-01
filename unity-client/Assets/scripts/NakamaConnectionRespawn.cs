using UnityEngine;
using Nakama;
using System;
using Nakama.TinyJson;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Linq;
using System.Threading;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class NakamaConnectionRespawn : MonoBehaviour
{
    public static NakamaConnectionRespawn Instance { get; private set; }

    [SerializeField] private string scheme = "http";
    [SerializeField] private string host = "127.0.0.1";
    [SerializeField] private int port = 7350;
    [SerializeField] private string serverKey = "defaultkey";
    
    [SerializeField] private Button startButton; // Reference to the UI button

    private IClient client;
    private ISession session;
    private ISocket socket;

    public IClient Client => client;
    public ISession Session => session;
    public ISocket Socket => socket;

    private void Awake()
    {
        // Ensure only one instance of this object
        if (Instance != null && Instance != this)
        {
            Debug.LogError("Duplicate NakamaConnection instance found! Destroying...");
            Destroy(gameObject);
        }
        Instance = this;

        if (startButton != null)
        {
            startButton.onClick.AddListener(StartNakamaConnection);
        }
        else
        {
            Debug.LogError("The button is not assigned in the script!");
        }
    }

    // Public method to start the connection when the button is clicked
    public async void StartNakamaConnection()
    {
        Debug.Log("Starting Nakama connection...");

        // 1. Create the client
        client = new Client(scheme, host, port, serverKey, UnityWebRequestAdapter.Instance);

        // 2. Authenticate the session
        session = await client.AuthenticateDeviceAsync(SystemInfo.deviceUniqueIdentifier);

        // 3. Create and connect the socket
        socket = client.NewSocket();
        await socket.ConnectAsync(session, true);

        Debug.Log($"Active session: {session}");
        Debug.Log($"Socket connected: {socket}");

        // 4. Check if the player already has a persona
        bool hasPersona = await CheckIfPersonaExists(session.UserId);
        if (!hasPersona)
        {
            try
            {
                string payload = $"{{\"personaTag\": \"{session.UserId[0..5]}\"}}";
                Debug.Log("Calling claim-persona RPC with payload: " + payload);
                var response = await client.RpcAsync(session, "nakama/claim-persona", payload);
                Debug.Log("Persona successfully registered: " + response.Payload);
            }
            catch (ApiResponseException ex)
            {
                Debug.LogError($"RPC Error: {ex.Message}");
            }
        }
        else
        {
            Debug.Log("Persona already registered!");
        }

        int maxRetries = 5; 
        
        for (int attempt = 1; attempt <= maxRetries; attempt++)
        {
            try
            {
                string payload = $"{{\"personaTag\": \"{session.UserId}\", \"nickname\": \"{session.UserId}\", \"display_name\": \"New-Player\"}}";
                Debug.Log($"Attempt {attempt}: Calling respawn-player RPC with payload: {payload}");
                
                var response = await client.RpcAsync(session, "tx/game/respawn-player", payload);
                Debug.Log("Player successfully created: " + response.Payload);
                // Wait a bit before we try again.
                for (float t = 0f; (t < 2f) && !Application.exitCancellationToken.IsCancellationRequested; t += Time.deltaTime)
                    await Task.Yield();
                SceneManager.LoadScene("MainScene");
                break;
            }
            catch (Exception ex)
            {
                Debug.LogError($"Attempt {attempt}/5 failed: {ex.Message}");
                
                if (attempt < maxRetries)
                {
                    Debug.Log($"Retrying in 0.5 seconds...");
                    // Wait a bit before we try again.
                    for (float t = 0f; (t < 0.5f) && !Application.exitCancellationToken.IsCancellationRequested; t += Time.deltaTime)
                        await Task.Yield();
                }
                else
                {
                    Debug.LogError("Max retry attempts reached. Unable to create player.");
                }
            }
        }
    }

    private async Task<bool> CheckIfPersonaExists(string userId)
    {
        try
        {
            string payload = "";
            var response = await client.RpcAsync(session, "nakama/show-persona", payload);
            Debug.Log("Persona found: " + response.Payload);
            return true;
        }
        catch (Exception ex)
        {
            Debug.LogError($"Persona not found: {ex.Message}");
            return false;
        }
    }
}