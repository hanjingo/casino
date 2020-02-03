using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ResManager  {

    private static ResManager instance;
    public static ResManager Instance
    {
        get
        {
            if (instance == null)
            {
                instance = new ResManager();
            }
            return instance;
        }
    }

    
}
