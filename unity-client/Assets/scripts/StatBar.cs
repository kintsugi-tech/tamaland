using UnityEngine;
using UnityEngine.UI;

public class StatBar : MonoBehaviour
{

    public Slider slider;

    public void SetMaxValue(int v)
    {
        slider.maxValue = v;
        slider.value = v;
    }
    
    public void SetValue(int v)
    {
        slider.value = v;
    }
}
