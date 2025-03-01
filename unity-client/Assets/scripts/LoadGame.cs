using UnityEngine;
using UnityEngine.UI;
using UnityEngine.SceneManagement;

public class LoadGame : MonoBehaviour
{
    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        // Get the Button component and add a listener to it
        GetComponent<Button>().onClick.AddListener(LoadGameData);
    }


    // Update is called once per frame
    void LoadGameData()
    {   
        // TODO: 1. load persona details ; 2. query game data
        
        SceneManager.LoadScene("MainScene");
    }
}   

