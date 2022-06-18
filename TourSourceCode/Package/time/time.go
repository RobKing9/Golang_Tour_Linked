package main

import (
	"fmt"
	"time"
)

func main() {
	//现在的时间
	now := time.Now()
	fmt.Println(now)
	//现在的年份
	fmt.Println(now.Year())
	//现在的月份
	fmt.Println(now.Month())
	//现在的天
	fmt.Println(now.Day())
	//现在的日期
	fmt.Println(now.Date())
	//现在的小时
	fmt.Println(now.Hour())
	//现在的分钟
	fmt.Println(now.Minute())
	//现在的秒
	fmt.Println(now.Second())
	fmt.Println("--------------")
	//睡眠的时间
	time.Sleep(2 * time.Second)
	//从之前到现在经过的时间
	fmt.Println(time.Since(now))
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	time := time.Unix(now.Unix(), 0)
	fmt.Println(time.Date())
	//时间格式化
	fmt.Println(now.Format("2006--01--02 3:04:05 PM"))
	//定时器
	Timer()
}

func Timer() {
	timer := time.NewTimer(3 * time.Second)
	//t1 := time.Now()
	//fmt.Println(t1)
	t2 := <-timer.C
	fmt.Println(t2)

	timer2 := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			<-timer2.C
			i++
			fmt.Println(time.Now())
			if i == 5 {
				timer2.Stop()
			}
		}
	}()
	for {
	}
}
