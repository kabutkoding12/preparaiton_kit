package main

import (
	"fmt"
	"github.com/kabutkoding12/preparaiton_kit/plusminus"
	"github.com/kabutkoding12/preparaiton_kit/minmaxsum"
	"github.com/kabutkoding12/preparaiton_kit/time_conversion"
	"github.com/kabutkoding12/preparaiton_kit/lonly_integer"
)

func main () {
	fmt.Println("=== plus minus ======")
	arr := []int32{2,5,4,0,-1,-2}
	plusminus.PlusMinus(arr)

	fmt.Println("=== min max sum ======")
	arr1 := []int{1, 2, 3, 4, 5}
	minmaxsum.MiniMaxSum(arr1)


	fmt.Println("=== time conversion ======")
	s := "07:15:45AM"
	fmt.Println(time_conversion.TimeConversion(s))

	fmt.Println("=== lonly integer ======")
	arrInteger := []int32{1,2,3,4,3,2,1}
	fmt.Println( lonly_integer.LonlyInteger(arrInteger) )


}