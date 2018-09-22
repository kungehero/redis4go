package main

import (
	"fmt"
	"redis4go"
)

func main() {
	p := redis4go.New("127.0.0.1:6379", "", 0, 100, 100)
	r := p.GetRedis()
	v := r.HSET("test", "1", "001")
	str := v.MustFloat64()
	fmt.Println(str)
	/*switch v.(type) {
	case string:
		fmt.Println(v)
	case int:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}*/
}
