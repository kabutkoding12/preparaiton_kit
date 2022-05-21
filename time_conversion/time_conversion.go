package time_conversion

import (
	"fmt"
	"strings"
	"strconv"
)

func TimeConversion(s string) string {

	s = strings.ToLower(s) //convert (AM or PM) to (am or pm) 
	pm := strings.HasSuffix(s, "pm") //check if suffix in string has pm and return value true or false

	am := strings.HasSuffix(s, "am") //check if suffix in string has am and return value true or false

	t := s[:len(s)-2] //cut two caracter the end of string
	
	timeArr := strings.Split(t, ":") //convert string to array with split index ":"
	
	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
	hhInt, err := strconv.Atoi(hh) //convert string to int
	

	if err != nil {
		panic(err.Error())
	}
	if am && hhInt == 12 {
		hhInt = 0
	}
	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt)	

	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)
	
}