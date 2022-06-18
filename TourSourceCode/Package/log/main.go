package main

import (
	"log/mylogger"
	"time"
)

func main () {
	mylog := mylogger.NewLog("debug")
	for {
		name := "zxw"
		age := 19
		mylog.Debug("这是一条debug日志 name:%s, age:%d", name, age)
		mylog.Info("这是一条info日志")
		mylog.Warning("这是一条warning日志")
		mylog.Error("这是一条error日志")
		mylog.Fatal("这是一条fatal日志")
		time.Sleep(time.Second)
	}
}
