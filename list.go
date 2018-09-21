/*list//数据结构 双向链表
lpush
rpush
lrange 获取有数据
linsert...after  在...之后插入数据
lindex 获取索引位置的数据

lpop rpop 非阻塞删除数据
blpop brpop 阻塞删除数据*/
package redis4go

//左端插入
func (this *Receiver) LPush(key string, args ...interface{}) *Result {
	return this.Do("LPUSH", key, args)
}

//右端插入
func (this *Receiver) RPush(key string, args ...interface{}) *Result {
	return this.Do("LPUSH", key, args)
}

//根据索引获取list
func (this *Receiver) Lrange(key string, startindex, endindex int) *Result {
	return this.Do("LPUSH", key, startindex, endindex)
}

//pivot (before/after) 在索引的中心点添加
func (this *Receiver) LInsert(key, pivot string, index int, value interface{}) *Result {
	return this.Do("LINSTER", key, pivot, index, value)

}

//获取索引位置处的值
func (this *Receiver) LIndex(key string, index int) *Result {
	return this.Do("LINDEX", key, index)

}

//非阻塞删除数据
func (this *Receiver) LPop(key string) *Result {
	return this.Do("LPOP", key)

}
func (this *Receiver) RPop(key string) *Result {
	return this.Do("RPOP", key)

}

//阻塞删除数据
func (this *Receiver) BLPop(key string) *Result {
	return this.Do("BLPOP", key)
}
func (this *Receiver) BRPop(key string) *Result {
	return this.Do("BRPOP", key)
}

//返回索引fanwei处的值，删除其他的值
func (this *Receiver) LTrim(key string, startindex, endindex int) *Result {
	return this.Do("LTRIM", key)
}

//设置索引处的值
func (this *Receiver) Lset(key string, index int) *Result {
	return this.Do("LSET", key, index)

}
