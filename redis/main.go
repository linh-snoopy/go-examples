package main

import (
	"log"
	//"reflect"
	//"bytes"
	"github.com/garyburd/redigo/redis"
	//"encoding/gob"
)

var (
	HashKey = "tse:test"
	TsePusher = "tsePusher"
)


type Test struct {
	name string
	age int
}

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
//	res, err := c.Do("HSET", HashKey,"foo", "bar")
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(res)
//	res, err = c.Do("HSET", HashKey,"never", "ever")
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(res)
	
//	str, err := redis.String(c.Do("HGET", HashKey, "foo"))
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(str)
	
//	strs, err := redis.Strings(c.Do("HVALS", HashKey))
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(strs)
	
//	t := Test{"Linh", 26}
//	buff := new(bytes.Buffer)
//	err = gob.NewEncoder(buff).Encode(t)
//	res, err := c.Do("HSET", HashKey, "test1", buff.Bytes())
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(res)
//	res, err = redis.String(c.Do("HGET", HashKey, "test1"))
//	log.Println(reflect.TypeOf(res))
//	if err != nil {
//		log.Println(err)
//	}
//	rbuff := new(bytes.Buffer)
//	err = gob.NewDecoder(rbuff).Decode(t)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println(t)
	c.Do("DEL", TsePusher)
	for i:= 0; i<10; i++ {
		res, err := c.Do("LPUSH", TsePusher, "a")
		if err != nil {
			log.Println(err)
		}
		log.Println(res)
	}
	res, err := c.Do("LRANGE", TsePusher, 0, -1)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}