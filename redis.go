package redis4go

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	errConn = errors.New("invalid Connect")
)

func New(addr, password string, dbIndex, maxActive, maxIdle int) (p *Pool) {
	var dialFunc = func() (c redis.Conn, err error) {
		if len(password) > 0 {
			c, err = redis.Dial("tcp", addr, redis.DialPassword(password))
		} else {
			c, err = redis.Dial("tcp", addr)
		}

		if err != nil {
			return nil, err
		}

		_, err = c.Do("SELECT", dbIndex)
		if err != nil {
			c.Close()
			return nil, err
		}

		return c, err
	}

	p = &Pool{}
	var pool = &redis.Pool{}
	pool.MaxIdle = maxIdle
	pool.MaxActive = maxActive
	pool.Wait = true
	pool.IdleTimeout = 180 * time.Second
	pool.Dial = dialFunc
	p.Pool = pool

	return p
}

func (this *Receiver) Do(CommdName string, args ...interface{}) *Result {
	if this.c != nil {
		return result(this.c.Do(CommdName, args))
	}
	return result(nil, errConn)
}

func (this *Receiver) Send(commandName string, args ...interface{}) *Result {
	var err = errConn
	if this.c != nil {
		err = this.c.Send(commandName, args...)
	}
	return result(nil, err)
}

func (this *Pool) GetRedis() *Receiver {
	var c = this.Pool.Get()
	if c == nil {
		return nil
	}
	return &Receiver{c: c}
}

type Pool struct {
	*redis.Pool
}

type Receiver struct {
	c Conn
}
type Conn interface {
	redis.Conn
}

func (this *Receiver) Transaction(callback func(conn Conn)) *Result {
	if this.c != nil {
		var c = this.c
		c.Send("MULTI")
		callback(c)
		return result(c.Do("EXEC"))
	}
	return result(nil, errConn)
}

func StructToArgs(key string, obj interface{}) redis.Args {
	return redis.Args{}.Add(key).AddFlat(obj)
}
