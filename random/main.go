//package main

//import (
//    "encoding/base64"
//    "crypto/rand"
//    "fmt"
// )

// func main() {
//   size := 32 // change the length of the generated random string here

//   rb := make([]byte,size)
//   _, err := rand.Read(rb)

//   if err != nil {
//      fmt.Println(err)
//   }

//   rs := base64.URLEncoding.EncodeToString(rb)

//   fmt.Println(rs)
// }
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(RandStringRunes(10))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
