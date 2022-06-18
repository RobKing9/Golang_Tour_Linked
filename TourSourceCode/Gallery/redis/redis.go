package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

type Client struct {
	name string
	psw  string
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	initClient()
	/*
		哈希
	*/
	var name, psw string
	var cli Client
	fmt.Scanln(&name)
	cli.name = name
	fmt.Scanln(&psw)
	cli.psw = psw
	rdb.HSet(cli.name, "name", cli.name)
	rdb.HSet(cli.name, "psw", cli.psw)
	name, err := rdb.HGet("zxw", "psw").Result()
	fmt.Println(psw)
	if err != nil {
		fmt.Println("hget failed:", err.Error())
	}
	//rdb.HMSet(cli.name, "name", cli.name, "psw", cli.psw)

	/*
			Set集合
		err := rdb.Set("score", 100, 0).Err()
		if err != nil {
			fmt.Printf("set score failed, err:%v\n", err)
			return
		}

		val, err := rdb.Get("score").Result()
		if err != nil {
			fmt.Printf("get score failed, err:%v\n", err)
			return
		}
		fmt.Println("score", val)

		val2, err := rdb.Get("name").Result()
		if err == redis.Nil {
			fmt.Println("name does not exist")
		} else if err != nil {
			fmt.Printf("get name failed, err:%v\n", err)
			return
		} else {
			fmt.Println("name", val2)
		}
	*/
}
