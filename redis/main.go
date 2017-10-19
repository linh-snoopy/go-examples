package main

import (
	"log"
	//"reflect"
	"bytes"
	"github.com/garyburd/redigo/redis"
	"encoding/gob"
	"strconv"
	"encoding/json"
)

var (
	HashKey = "tse:test"
	TsePusher = "tsePusher"
)


type Test struct {
	Name string `json:"name"`
	Age int `json:"age"`
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
	var t Test
	c.Do("DEL", TsePusher)
	r3, err := redis.String(c.Do("lpop", TsePusher))
	if err != nil {
		log.Println(err)
		if ok, _ := checkQueueEmpty(c); ok {
			log.Println("Empty DB")
		} else {
			return
		}
	}
	inn := []byte(r3)
	err = json.Unmarshal(inn, &t)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(t)
	for i:= 0; i<10; i++ {
		test := Test{
			Name: "Linh_" + strconv.Itoa(i),
			Age: 26,
		}
		bTest, err := json.Marshal(test)
		if err != nil {
			log.Println(err)
			return
		}
		res, err := c.Do("LPUSH", TsePusher, bTest)
		if err != nil {
			log.Println(err)
		}
		log.Println(res)
	}
	res, err := redis.Strings(c.Do("LRANGE", TsePusher, 0, -1))
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	r, err := redis.Strings(c.Do("blpop", TsePusher, 0))
	if err != nil {
		log.Println(err)
	}
	in := []byte(r[1])
	err = json.Unmarshal(in, &t)
	if err != nil {
		return
	}
	log.Println(t)
	r2, err := redis.String(c.Do("lpop", TsePusher))
	if err != nil {
		log.Println(err)
	}
	in = []byte(r2)
	err = json.Unmarshal(in, &t)
	if err != nil {
		return
	}
	log.Println(t)
}

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getInterface(bts []byte, data interface{}) error {
	buf := bytes.NewBuffer(bts)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func checkQueueEmpty(c redis.Conn) (bool, error){
	res, err := redis.Int(c.Do("LLEN", TsePusher))
	if err != nil {
		return false, nil
	}
	log.Println(res)
	if res > 0 {
		return false, nil
	}
	return true, nil
}