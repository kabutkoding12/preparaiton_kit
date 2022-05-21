package lonly_integer

// import "fmt"


func LonlyInteger(a []int32) int32{

	var result int32

	for i:=0; i<len(a); i++ {
		result ^= a[i]
	}

	return result
}