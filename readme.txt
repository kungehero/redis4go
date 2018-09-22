sttring

set添加数据 ok
get获取数据 nil
strlrn 获取key值的长度 0
append直接添加 返回添加之后的长度
setrange 索引添加 同上

setnx 原子性添加数据    true 1 false 0
setxx  键值存在时设置
mset 设置多个
mget 获取多个

list//数据结构 双向链表
lpush 
rpush
lrange 获取有数据
linsert...after  在...之后插入数据
lindex 获取索引位置的数据

lpop rpop 非阻塞删除数据
blpop brpop 阻塞删除数据

hash
//map或字典 结构  最多容纳2*32-1个字段                                 

hmset
hmget

hget 获取单个值
hexists判断是否存在字段  返回1或0
hgetall 获取 一个hash中所有的字段和值
hdel 删除哈希字段
hsetnx  字段不存在的情况下设置

hscan 增量获取所用字段和值 
如：hscan hash块 0（所以） math （字符串） count 10（默认值吗）

set 集合(唯一，无序对象的组合)   hashtable 用于数据的存在 删除  组合
2*23-1
无序集合：
sadd 一键多值
sismember 测试存在
srem 删除哈希字段
scard  获取成员数量
smembers 获取所有几何元素
sscan 获取增量元素

sunion/sunionstore 并集
sinter/sinterstore 交集
sdiff/sdiffstore 差集

有序集合：
zadd 添加数据  zaddnx添加新数据
zrevrange 获取集合数据（排序）
zincrby 修改数据
zrevrank 获取排名
zscore 获取投票数


hyperloglog(Hll)  适用统计个数

pfadd 添加数据 
pfcount 统计个数
pfmerge   合并
Geo(gis地图  设计) 适用坐标统计

geoadd 添加坐标
geopos  获取指定坐标
georadius   获取指定范围内的坐标
geodist   获取坐标间的距离

键管理

dbsize  获取键数量
keys * 获取所有的键  scan 增量获取键值

del/unlink 删除键
exists 判断键是否存在

type  获取键的数据类型（如：hash,string……）

