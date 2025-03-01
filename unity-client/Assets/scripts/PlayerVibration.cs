using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.InputSystem;


public class PlayerVibration : MonoBehaviour
{

    public float vibrationInterval = 10.0f;
    public float vibrationDuration = 0.25f; // How long the vibration lasts
    public float vibrationAmount = 0.5f;
    public float vibrationSpeed = 50f; // Speed of vibration
    private float vibrationTimer = 0.0f;
    private Vector3 originalPosition;
    private bool isVibrating = false;
    private float nextVibrationTime = 0f;
    

    void Start()
    {
        originalPosition = transform.position;
        nextVibrationTime = Time.time + vibrationInterval;
    }

    // Update is called once per frame
    void Update()
    {
        // Check if it's time to start vibrating
        if (Time.time >= nextVibrationTime && !isVibrating)
        {
            isVibrating = true;
            vibrationTimer = 0f;
        }
        // Handle vibration
        if (isVibrating)
        {
            if (vibrationTimer < vibrationDuration)
            {
                float offset = Mathf.Sin(Time.time * vibrationSpeed) * vibrationAmount;
                transform.position = new Vector3(originalPosition.x + offset, originalPosition.y, originalPosition.z);

                vibrationTimer += Time.deltaTime;
            }
            else
            {
                // Stop vibrating and reset position
                isVibrating = false;
                transform.position = originalPosition;
                nextVibrationTime = Time.time + vibrationInterval; // Schedule next vibration
            }
        }
    }

}
