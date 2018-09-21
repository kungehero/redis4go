/*hash
//map或字典 结构  最多容纳2*32-1个字段
hmset
hmget添加单个字段值
hset
hget 获取单个值
hexists判断是否存在字段  返回1或0
hgetall 获取 一个hash中所有的字段和值
hdel 删除哈希字段
hsetnx  字段不存在的情况下设置
hscan 增量获取所用字段和值
如：hscan hash块 0（所以） math （字符串） count 10（默认值吗）*/
package redis4go

//设置多个字段值
func (this *Receiver) HMSet(key string, args ...interface{}) *Result {
	var ps []interface{}
	ps = append(ps, key)
	if len(args) > 0 {
		ps = append(ps, args...)
	}
	return this.Do("HMSET", ps...)
}

//获取指定自断的值
func (this *Receiver) HMGet(key string, fields ...interface{}) *Result {
	var ks []interface{}
	ks = append(ks, key)
	for _, f := range fields {
		ks = append(ks, f)
	}
	return this.Do("HMGET", ks...)
}

//删除hash表正宗指定字段的值
func (this *Receiver) HDel(key string, fields ...string) *Result {
	var ks []interface{}
	ks = append(ks, key)
	for _, f := range fields {
		ks = append(ks, f)
	}
	return this.Do("HDEL", ks...)
}

//HEXISTS 查看哈希表 key 中，给定域 field 是否存在。
func (this *Receiver) HEXISTS(key, field string) *Result {
	return this.Do("HEXISTS", key, field)
}

//HGET 返回哈希表 key 中给定域 field 的值。
func (this *Receiver) HGET(key string, field string) *Result {
	return this.Do("HGET", key, field)
}

// HGETALL 返回哈希表 key 中，所有的域和值。
func (this *Receiver) HGETALL(key string) *Result {
	return this.Do("HGETALL", key)
}

//HINCRBY 为哈希表 key 中的域 field 的值加上增量 increment 。
func (this *Receiver) HINCRBY(key, field string, increment int) *Result {
	return this.Do("HINCRBY", key, field, increment)
}

//HINCRBYFLOAT 为哈希表 key 中的域 field 加上浮点数增量 increment 。
func (this *Receiver) HINCRBYFLOAT(key, field string, increment float64) *Result {
	return this.Do("HINCRBYFLOAT", key, field, increment)
}

//HKEYS 返回哈希表 key 中的所有域。
func (this *Receiver) HKEYS(key string) *Result {
	return this.Do("HKEYS", key)
}

//HLEN 返回哈希表 key 中域的数量。
func (this *Receiver) HLEN(key string) *Result {
	return this.Do("HLEN", key)
}

func (this *Receiver) HMSETStruct(key string, obj interface{}) *Result {
	return this.Do("HMSET", StructToArgs(key, obj)...)
}

//HSET 将哈希表 key 中的域 field 的值设为 value 。
func (this *Receiver) HSET(key, field string, value interface{}) *Result {
	return this.Do("HSET", key, field, value)
}

//HSETNX 将哈希表 key 中的域 field 的值设置为 value ，当且仅当域 field 不存在。
func (this *Receiver) HSETNX(key, field string, value interface{}) *Result {
	return this.Do("HSETNX", key, field, value)
}

//HVALS 返回哈希表 key 中所有域的值。
func (this *Receiver) HVALS(key string) *Result {
	return this.Do("HVALS", key)
}

// HSTRLEN 返回哈希表 key 中， 与给定域 field 相关联的值的字符串长度（string length）。
func (this *Receiver) HSTRLEN(key, field string) *Result {
	return this.Do("HSTRLEN", key, field)
}
