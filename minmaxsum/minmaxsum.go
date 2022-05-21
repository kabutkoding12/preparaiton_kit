package minmaxsum

import(
	"fmt"
)



func MiniMaxSum(arr []int) {
	// Write your code here
	
	var sumValue int

	for i:=0; i<len(arr); i++ {
		sumValue += arr[i]
	}

	var min, max int

	for i:=0; i<len(arr); i++ {

		if sumValue - arr[i] <= min || min == 0 {
			min = sumValue - arr[i]
		}
		if sumValue - arr[i] > max || max == 0 {
			max = sumValue - arr[i]
		}
	}

	fmt.Printf("%d %d \n", min, max)
}