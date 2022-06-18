package main

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"strconv"
)

func main() {
	id := ksuid.New()
	fmt.Println(id.String())
	userid, err := strconv.ParseInt(id.String(), 10, 64)
	if err != nil {
		fmt.Println("parse failed:", err.Error())
	}
	fmt.Println(userid)

	getXid()
}
func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}
