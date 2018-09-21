package redis4go

func Test() {
	println("ceshi")
}

//返回ok
func (this *Receiver) Set(params ...interface{}) *Result {
	return this.Do("SET", params)
}

//返回key值  字符串
func (this *Receiver) Get(key string) *Result {
	return this.Do("GET", key)
}

//返回字符串长度  不存在返回0
func (this *Receiver) Strlrn(key string) *Result {
	return this.Do("STRLEN", key)
}

//像一个键的字符串末尾添加
func (this *Receiver) Append(key, value string) *Result {
	return this.Do("APPEND", key, value)
}

//返回修改后的字符串
func (this *Receiver) SetRange(key string, strindex, endindex int) *Result {
	return this.Do("SETRANGE", key, strindex, endindex)
}

//键存在返回0 不添加    返回1 添加键值对
func (this *Receiver) SetNx(key, value string) *Result {
	return this.Do("SETNX", key, value)
}

//批量添加键值对
func (this *Receiver) MSet(args interface{}) *Result {
	return this.Do("MSET", args)
}

//批量获取键值
func (this *Receiver) MGet(key ...string) *Result {

	var keys []interface{}
	for _, k := range key {
		keys = append(keys, k)
	}
	return this.Do("MSET", keys)
}
