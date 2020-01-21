package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("  -78"))
	fmt.Println(int('0'))
	fmt.Println(byte(48 - '0'))
}

func myAtoi(str string) int {
	res, sign, l, idx := 0, 1, len(str), 0
	for idx < l && (str[idx] == ' ' || str[idx] == '\t') {
		idx++
	}

	if idx < l {
		if str[idx] == '+' {
			sign = 1
			idx++
		} else if str[idx] == '-' {
			sign = -1
			idx++
		}
	}

	for idx < l && str[idx] >= '0' && str[idx] <= '9' {
		res = res*10 + int(str[idx]) - '0'
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if res*sign < math.MinInt32 {
			return math.MinInt32
		}
		idx++
	}

	return res * sign

	/*
	   res, sign, len, idx := 0, 1, len(str), 0

	   // Skip leading spaces
	   for idx < len && (str[idx] == ' ' || str[idx] == '\t') {
	       idx++
	   }

	   // +/- Sign
	   if idx < len {
	       if str[idx] == '+' {
	           sign = 1
	           idx++
	       } else if str[idx] == '-' {
	           sign = -1
	           idx++
	       }
	   }

	   // Numbers
	   for idx < len && str[idx] >= '0' && str[idx] <= '9'{
	       res = res * 10 + int(str[idx]) - '0'
	       if sign * res > math.MaxInt32 {
	           return math.MaxInt32
	       } else if sign * res < math.MinInt32 {
	           return math.MinInt32
	       }
	       idx++
	   }

	   return res * sign
	*/
}
