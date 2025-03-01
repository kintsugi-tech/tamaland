using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using Newtonsoft.Json; 
using Nakama;
using System.Threading.Tasks; 
using Newtonsoft.Json;
using UnityEngine.SceneManagement;
using UnityEngine.UI;


public class PlayerController : MonoBehaviour
{

private struct State
{
    [JsonProperty("State")]
    public string state;

    [JsonProperty("EndStateTimestamp")] 
    public long endStateTimestamp;
}

private struct LastUpdate
{
    [JsonProperty("Tick")] 
    public int tick;

    [JsonProperty("Timestamp")]
    public long timestamp; 
}

private struct PlayerStatsResponse
{
    [JsonProperty("H")]
    public int health;

    [JsonProperty("F")]
    public int food;

    [JsonProperty("E")]
    public int energy;

    [JsonProperty("L")]
    public int level;

    [JsonProperty("S")]
    public State state;

    [JsonProperty("LU")]
    public LastUpdate lastUpdate;
}


    public int maxHealth = 100;
    public int maxFood = 100;
    public int maxEnergy = 100;
    public int currentHealth;
    public int currentFood;
    public int currentEnergy;
    public int currentLevel;

    public StatBar healthBar;
    public StatBar foodBar;
    public StatBar energyBar;
    public string currentState;
    public Button eatButton;
    public Button sleepButton;
    public Text levelText;


    public int minHealth = 0;
    public int minFood = 0;
    public int minEnergy = 0;

    private async void Start()
    {
        Debug.Log("Fetching player stats...");

        PlayerStatsResponse response = await PlayerStats(NakamaConnection.Instance.Session.UserId);

        currentHealth = response.health;
        currentFood = response.food;
        currentEnergy = response.energy;
        currentState = response.state.state;
        currentLevel = response.level;
        healthBar.SetMaxValue(maxHealth);
        foodBar.SetMaxValue(maxFood);
        energyBar.SetMaxValue(maxEnergy);
        levelText.text = UpdateLevelText(currentLevel);
    }

    private string UpdateLevelText(int lv) {
        return "Level: " + lv;
    }

    private async void Update()
    {
        PlayerStatsResponse response = await PlayerStats(NakamaConnection.Instance.Session.UserId);
        Debug.Log($"response: {response}");
        Debug.Log($"response.health: {response.health}");
        Debug.Log($"response.level: {response.level}");

        currentHealth = response.health;
        currentFood = response.food;
        currentEnergy = response.energy;
        currentState = response.state.state;
        currentLevel = response.level;
        if (currentState == "Dead")
        {
            // Handle death logic here
            Debug.Log("Player is dead!");
            SceneManager.LoadScene("Dead");
        }

        healthBar.SetValue(currentHealth);
        foodBar.SetValue(currentFood);
        energyBar.SetValue(currentEnergy);

        if (currentState=="Dead" || currentState=="Sleeping" || currentState=="Eating") {
            eatButton.interactable = false; 
            sleepButton.interactable = false;
        } else {
            eatButton.interactable = true; 
            sleepButton.interactable = true;
        }

        levelText.text = UpdateLevelText(currentLevel);

        Debug.Log($"Health: {currentHealth}/{maxHealth}, Food: {currentFood}/{maxFood}, Energy: {currentEnergy}/{maxEnergy}, State: {currentState}");
    }

    private async Task<PlayerStatsResponse> PlayerStats(string nickname)
    {
        try
        {
            string payload = $"{{\"nickname\": \"{nickname}\"}}";
            Debug.Log("Calling player-stats RPC with payload: " + payload);

            var response = await NakamaConnection.Instance.Client.RpcAsync(
                NakamaConnection.Instance.Session,
                "query/game/player-stats",
                payload
            );

            Debug.Log("Received player stats: " + response.Payload);
            return JsonConvert.DeserializeObject<PlayerStatsResponse>(response.Payload);
        }
        catch (ApiResponseException ex)
        {
            Debug.LogError($"RPC Error: {ex.Message}");
            return new PlayerStatsResponse(); 
        }
    }
}
