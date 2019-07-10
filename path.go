package main

import "fmt"

func main() {
	// dir , err := os.Getwd()
	//dir , err := filepath.Abs("login.gtpl")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//mt.Pritln(dir)

	// pase_student()

	// println(DeferFunc1(1))
	// println(DeferFunc2(1))
	// println(DeferFunc3(1))

	fmt.Println(x,y,z,k,p)
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu  // 指向数组的最后一个
	}

	for k,v:=range m{
		println(k,"=>",v.Name)
	}

	// 正确
	for i:=0;i<len(stus);i++  {
		m[stus[i].Name] = &stus[i]
	}
	for k,v:=range m{
		println(k,"=>",v.Name)
	}

}

type student struct {
	Name string
	Age  int
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
