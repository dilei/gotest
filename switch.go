package main

import "fmt"

func main() {
	arr64 := []int64{1,2,3,4}
	for _, val := range arr64 {
		switch val {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)

		default:
			fmt.Println(4)
			// strArr := [9]int{1,2,3,4,5,6,7,8,9}
			// strSlice := []int{1,2,3,4,5,6,7,8,9}
			strMap := make(map[int]string)
			strMap[1] = "1"
			strMap[2] = "2"
			strMap[3] = "3"
			strMap[4] = "4"
			strMap[5] = "5"
			strMap[6] = "6"
			strMap[7] = "7"
			strMap[8] = "8"
			strMap[9] = "9"



			for i:=0; i<1000; i++ {
				// fmt.Println(strArr)
				// fmt.Println(strSlice)
				fmt.Println(strMap)
			}
		}
	}
}
