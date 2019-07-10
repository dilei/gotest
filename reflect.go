package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Username string
	age int
}

type Admin struct {
	User
	title string
}

type User1 struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age int `field:"age" type:"tinyint"`
}

func main() {
	var user Admin
	t := reflect.TypeOf(user)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	f, _ := t.FieldByName("title")
	fmt.Println(f.Name)

	f, _ = t.FieldByName("Username") // 直接访问嵌⼊入字段成员，会⾃自动深度查找。
	fmt.Println(f.Name)

	f = t.FieldByIndex([]int{0, 1}) // Admin[0] -> User[1] -> age
	fmt.Println(f.Name)

	// 元数据编程
	var u1 User1
	t1 := reflect.TypeOf(u1)
	f1, _ := t1.FieldByName("Name")
	fmt.Println(f1.Tag)
	fmt.Println(f1.Tag.Get("field"))
	fmt.Println(f1.Tag.Get("type"))

	// Value

	admin := &Admin{User{"name", 12}, "NT"}
	v := reflect.ValueOf(admin).Elem()

	fmt.Println(v.FieldByName("Username").String())
	fmt.Println(v.FieldByName("title").String())
	fmt.Println(v.FieldByIndex([]int{0,1}).Int())

	// 除具体的 Int、String 等转换⽅方法，还可返回 interface{}。只是⾮非导出字段⽆无法使⽤用，需
	// ⽤ CanInterface 判断⼀一下
	u := User{"Jack", 23}
	v = reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Username").Interface())
	// fmt.Println(v.FieldByName("age").Interface())  // panic: reflect.Value.Interface: cannot return value obtained from unexported field or method

	// 转具体类型可以
	f2 := v.FieldByName("age")
	if f2.CanInterface() {
		fmt.Println(f2.Interface())
	} else {
		fmt.Println(f2.Int())
	}

	// slice 取值
	vSlice := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, vSlice.Len(); i < n; i++ {
		fmt.Println(vSlice.Index(i).Int())
	}

	// map 取值
	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}

	// 需要注意，Value 某些⽅方法没有遵循 "comma ok" 模式，⽽而是返回 ZeroValue，因此需
	// 要⽤用 IsValid 判断⼀一下是否可⽤用
	u = User{}
	v2 := reflect.ValueOf(u)
	f2 = v2.FieldByName("a")
	fmt.Println(f2.Kind(), f2.IsValid())


	// 将对象转换为接⼝口，会发⽣生复制⾏行为。该复制品只读，⽆无法被修改。所以要通过接⼝口改变
	//⽬目标对象状态，必须是 pointer-interface。
	//就算是指针，我们依然没法将这个存储在 data 的指针指向其他对象，只能透过它修改⺫⽬目
	//标对象。因为⽬目标对象并没有被复制，被复制的只是指针。
	u = User{"Jack", 23}
	v = reflect.ValueOf(u)
	p := reflect.ValueOf(&u)
	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())

	//⾮非导出字段⽆无法直接修改，可改⽤用指针操作。
	u = User{"Jack", 23}
	p = reflect.ValueOf(&u).Elem()
	p.FieldByName("Username").SetString("Tom")
	f3 := p.FieldByName("age")
	fmt.Println(f3.CanSet())
	// 判断是否能获取地址。
	if f3.CanAddr() {
		age := (*int)(unsafe.Pointer(f3.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}
	// 注意 p 是 Value 类型，需要还原成接⼝口才能转型。
	fmt.Println(u, p.Interface().(User))
}
