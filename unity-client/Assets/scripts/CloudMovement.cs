using UnityEngine;

public class CloudMovement : MonoBehaviour
{
    public float speed = 2.0f;
    public float xDestroyTrigger = -463.0f;

    // Start is called once before the first execution of Update after the MonoBehaviour is created
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        transform.position += Vector3.left * speed * Time.deltaTime;

        if (transform.position.x < xDestroyTrigger) {
            Destroy(gameObject);
        }
    }
}
