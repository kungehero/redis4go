package redis4go

/*pfadd 添加数据
pfcount 统计个数
pfmerge   合并

Geo(gis地图  设计) 适用坐标统计
geoadd 添加坐标
geopos  获取指定坐标
georadius   获取指定范围内的坐标
geodist   获取坐标间的距离*/

func (this *Receiver) PfAdd(key string, value interface{}) *Result {
	return this.Do("PFADD", key, value)
}
func (this *Receiver) Pfcount(key string) *Result {
	return this.Do("PFCOUNT", key)
}
func (this *Receiver) Pfmerge(newkey string, sourcekey ...interface{}) *Result {
	return this.Do("PFMERGE", newkey, sourcekey)
}

/*
func (this *Receiver) Geo() *Result {
	return this.Do("GEO")
}
func (this *Receiver) Geoadd() *Result {
	return this.Do("GEOADD")
}
func (this *Receiver) Geopos() *Result {
	return this.Do("GEOPOS")
}
func (this *Receiver) Georadius() *Result {
	return this.Do("GEORADIUS")
}
func (this *Receiver) Geodist() *Result {
	return this.Do("GEODIST")
}*/
