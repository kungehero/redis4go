package main

import (
	"fmt"
	"redis4go"
)

func main() {
	p := redis4go.New("127.0.0.1:6379", "", 0, 100, 100)
	r := p.GetRedis()
	v := r.Get("922name")
	str := v.MustString()
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
