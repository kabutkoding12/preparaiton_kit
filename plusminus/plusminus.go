package plusminus

import (
	"fmt"
)

func PlusMinus(arr []int32) {
    // Write your code here
    n  := len(arr)
    var plus, minus, zero int;
    
    for i :=0; i < n; i++ {
        
        if arr[i] == 0{
            zero++;
        }else if arr[i] < 0{
            minus++;
        }else if arr[i] > 0 {
            plus++
        }
        
    }

    plusValue := float64(plus) / float64(n)
    minusValue := float64(minus) / float64(n)
    zeroValue := float64(zero) / float64(n)
    
    fmt.Printf("%.6f\n",plusValue)
    fmt.Printf("%.6f\n",minusValue)
    fmt.Printf("%.6f\n",zeroValue)
    

}


