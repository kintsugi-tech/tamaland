using UnityEngine;

public class CloudSpawnerScript : MonoBehaviour
{

    public GameObject cloudPrefab;
    public float spawnRate = 400.0f;
    private float timer = 0.0f;
    public float spawnYRandLimitPercent = 5.0f;

    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        Debug.Log("transform.position.y: " + transform.position.y);
        Debug.Log("ScreenHeight: " + Screen.currentResolution.height);
        spawnCloud();
    }

    // Update is called once per frame
    void Update()
    {   
        if (timer < spawnRate) {
            timer += Time.deltaTime;
        } else {
            timer = 0.0f;
            spawnCloud();
        }
    }

    void spawnCloud() {
        float screenHeightPercent = Screen.currentResolution.height/100;
        float increment = screenHeightPercent*Random.Range(-spawnYRandLimitPercent, spawnYRandLimitPercent);
        float newYPosition = transform.position.y + increment;
        Vector3 spawnPosition = new Vector3(transform.position.x, newYPosition, transform.position.z);
        Debug.Log("New Y Position: " + newYPosition + " | increment: " + increment + " screenHeightPercent: " + screenHeightPercent);
        
        GameObject newCloud = Instantiate(cloudPrefab, spawnPosition, transform.rotation);
        newCloud.transform.localScale = new Vector3(2 * cloudPrefab.transform.localScale.x, 2 * cloudPrefab.transform.localScale.y, cloudPrefab.transform.localScale.z);
    }
}
