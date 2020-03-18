package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Print("conn error")
	}
	c.Do("SELECT", "10")
	c.Do("SET", "hello", "world")
	res, err := c.Do("GET", "hello")
	s, err := redis.String(res, err)
	fmt.Printf("%+v", s)
	sub(c)
	defer c.Close()
}

func sub(c redis.Conn) error {
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe("example")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return v
		}
	}
}
