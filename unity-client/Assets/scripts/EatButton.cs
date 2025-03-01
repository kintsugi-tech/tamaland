using UnityEngine;
using UnityEngine.UI;
using System;
public class FoodButton : MonoBehaviour
{
    public PlayerController player;  // Drag the player GameObject here in the Inspector

    void Start()
    {
        // Get the Button component and add a listener to it
        GetComponent<Button>().onClick.AddListener(RestoreFood);
    }

    async void RestoreFood()
    {
        if (player != null)
        {
            try
            {
                string payload = $"{{\"target\": \"{NakamaConnection.Instance.Session.UserId}\"}}";
                Debug.Log("Calling eat RPC with payload: " + payload);

                var response = await NakamaConnection.Instance.Client.RpcAsync(
                    NakamaConnection.Instance.Session,
                    "tx/game/eat",
                    payload
                );

                Debug.Log("Player successfully eat: " + response.Payload);
            }
            catch (Exception ex)
            {
                Debug.LogError($"RPC Error: {ex.Message}");
            }
        }
    }
}