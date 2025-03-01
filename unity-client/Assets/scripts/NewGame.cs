using UnityEngine;
using UnityEngine.UI;
using UnityEngine.SceneManagement;

public class NewGame : MonoBehaviour
{
    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        // Get the Button component and add a listener to it
        GetComponent<Button>().onClick.AddListener(InitializeGame);
    }


    // Update is called once per frame
    void InitializeGame()
    {   
        // TODO: 1. Create Account ; 2. Create Character

        // TODO: save persona details
        
        SceneManager.LoadScene("MainScene");
    }
}   

