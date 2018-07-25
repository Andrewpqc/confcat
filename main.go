package main

import (
	"fmt"
	"time"
	"confcat/config"
	"log"
)

func main() {
	c, err := config.NewConfig("test.conf")
	if err != nil {
		log.Fatal("error to new config:%v", err)
	}

	for {
		fmt.Println(c.GetString("host", "127.0.0.1"))
		fmt.Println(c.GetInt("port", 235))
		fmt.Println(c.GetFloat("P", 3.14))
		time.Sleep(3 * time.Second)
	}
}
