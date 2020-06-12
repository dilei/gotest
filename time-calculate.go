package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间:", now)
	//两小时之后的时间
	h, _ := time.ParseDuration("2h")
	nowAfter2Hour := now.Add(h)
	fmt.Println("两小时之后的时间:", nowAfter2Hour)

	//两小时之前的时间
	negativeH, _ := time.ParseDuration("-2h")
	nowBefore2Hour := now.Add(negativeH)
	fmt.Println("两小时之前的时间:", nowBefore2Hour)

	//十分钟之后的时间
	m, _ := time.ParseDuration("10m")
	nowAfter10Minute := now.Add(m)
	fmt.Println("十分钟之后的时间:", nowAfter10Minute)

	//十分钟之前的时间
	negativeM, _ := time.ParseDuration("-10m")
	nowBefore10Minute := now.Add(negativeM)
	fmt.Println("十分钟之前的时间:", nowBefore10Minute)

	//30s之后的时间
	s, _ := time.ParseDuration("30s")
	nowAfter30Second := now.Add(s)
	fmt.Println("30s之后的时间:", nowAfter30Second)

	//30s之前的时间
	negativeS, _ := time.ParseDuration("-30s")
	nowBefore30Second := now.Add(negativeS)
	fmt.Println("30s之前的时间:", nowBefore30Second)

	//1年1个月1天 之后的时间
	fmt.Println("1年2个月3天之后的时间:", now.AddDate(1, 2, 3))

	//2年2个月2天之前的时间
	fmt.Println("2年3个月4天之前的时间:", now.AddDate(-2, -3, -4))

	// Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t)
	t1 := time.Now()
	time.Sleep(2 * time.Second)
	t2 := time.Since(t1)
	t3 := time.Now().Sub(t1)

	fmt.Println(t2, t3, t3.String())
	fmt.Println(t2.Seconds())
	// 10s
	fmt.Print(time.Duration(10) * time.Second)
}
