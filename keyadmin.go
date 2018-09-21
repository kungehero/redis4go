//键 管理

/*dbsize  获取键数量
keys * 获取所有的键  scan 增量获取键值
del/unlink 删除键
exists 判断键是否存在
type  获取键的数据类型（如：hash,string……）*/
package redis4go

func (this *Receiver) DBSize() *Result {
	return this.Do("DBSIZE")
}

func (this *Receiver) Keys(pattern string) *Result {
	return this.Do("KEYS", pattern)
}
func (this *Receiver) Delete(key string) *Result {
	return this.Do("DEL", key)
}

//在redis.4.0以上版本引入，主要执行大KEY的异步删除
func (this *Receiver) UNLink(key string) *Result {
	return this.Do("UNLINK", key)
}
func (this *Receiver) Exists(key string) *Result {
	return this.Do("EXiSTS", key)
}
func (this *Receiver) Type(key string) *Result {
	return this.Do("TYPE", key)
}
