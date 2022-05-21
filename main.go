package main

import (
	"fmt"
	"github.com/kabutkoding12/preparaiton_kit/plusminus"
	"github.com/kabutkoding12/preparaiton_kit/minmaxsum"
)

func main () {
	fmt.Println("=== plus minus ======")
	arr := []int32{2,5,4,0,-1,-2}
	plusminus.PlusMinus(arr)

	fmt.Println("=== min max sum ======")
	arr1 := []int{1, 2, 3, 4, 5}
	minmaxsum.MiniMaxSum(arr1)


	fmt.Println("=== tess ======")

}