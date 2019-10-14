package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	var P person // P现在就是person类型的变量了

	P.name = "Astaxie"                            // 赋值"Astaxie"给P的name属性.
	P.age = 25                                    // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s", P.name) // 访问P的name属性.

	type Skills []string

	type Human struct {
		name   string
		age    int
		weight int
	}

	type Student struct {
		Human      // 匿名字段，struct
		Skills     // 匿名字段，自定义的类型string slice
		int        // 内置类型作为匿名字段
		speciality string
	}

	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)

}
