package main

import "fmt"

func main() {
	var arr1 []int
	arr1 = append(arr1, []int{1, 2, 3, 4}...)
	fmt.Println(arr1, len(arr1), cap(arr1))
	// arr1 = arr1[0:5] // 下标不能超过cap

	arr2 := make([]int, len(arr1), cap(arr1)*2)
	copy(arr2, arr1)
	arr1 = arr2
	fmt.Println(arr1, len(arr1), cap(arr1))

	arr1 = arr1[0:5]
	arr1[4] = 5
	fmt.Println(arr1, len(arr1), cap(arr1))

	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	slice := data[1:4:5] // [low : high : max]
	fmt.Println(slice, len(slice), cap(slice))

	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使⽤用索引号。
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8) // 使⽤用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))

	ageSlice := []sdog{
		{Name: "dd", Age: 12},
		{Name: "dd2", Age: 13},
	}
	fmt.Println(ageSlice)
	for key, val := range ageSlice {
		val.Age = val.Age + 1
		ageSlice[key] = val
	}
	fmt.Println(ageSlice)

	var str map[int]int = nil
	fmt.Println(len(str))

	// 合并切片
	var nullArr []int
	var arr = []int{1, 2, 3}
	arr3 := append(nullArr, arr...)
	fmt.Println(arr3)

	// 遍历空切片
	for key, val := range nullArr {
		fmt.Println("null slice")
		fmt.Println(key, val)
	}

	ss := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(ss), cap(ss))
	// fmt.Println(remove(ss, 5))
	// fmt.Println(ss)
	fmt.Println(remove(ss, 0))
	fmt.Println(len(ss), cap(ss))
	fmt.Println(ss)

}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

type sdog struct {
	Name string
	Age  int
}
