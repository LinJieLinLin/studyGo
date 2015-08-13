package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now().UnixNano()
	len := 1000000
	s := []string{}
	for i := 0; i < len; i++ {
		s = append(s, `纳秒，时间单位。一秒的10亿分之一，即等于10的负9次方秒。常用作内存读写速度的单位，其前面数字越小则表示速度越快。
					纳秒也是计算机中的一个专业术语，是用来计算CPU及各个硬件所运行的速度的运行单位。
					1纳秒=0.000000001秒=10^(-9)秒`)
	}
	fmt.Println("1", float64((time.Now().UnixNano()-t1))/float64(int64(time.Second)), "秒\n")
	fmt.Println("now len:",len(s))
	one(s)
	two(s)
	fmt.Println("len 在 1000000 以内，_,o:=range 比i,_:=range 好")
	for i := 0; i < len*9; i++ {
		s = append(s, `纳秒，时间单位。一秒的10亿分之一，即等于10的负9次方秒。常用作内存读写速度的单位，其前面数字越小则表示速度越快。
					纳秒也是计算机中的一个专业术语，是用来计算CPU及各个硬件所运行的速度的运行单位。
					1纳秒=0.000000001秒=10^(-9)秒`)
	}
	fmt.Println("1", float64((time.Now().UnixNano()-t1))/float64(int64(time.Second)), "秒\n")
	fmt.Println("now len:",len(s))
	one(s)
	two(s)
}
func one(arg_s []string) {
	t1 := time.Now().UnixNano()
	for i, _ := range arg_s {
		arg_s[i] = arg_s[i] + "123"
	}
	fmt.Println("2", float64((time.Now().UnixNano()-t1))/float64(int64(time.Second)), "秒\n")
}
func two(arg_s []string) {
	t1 := time.Now().UnixNano()
	for _, o := range arg_s {
		o = o + "123"
	}
	fmt.Println("3", float64((time.Now().UnixNano()-t1))/float64(int64(time.Second)), "秒\n")
}
