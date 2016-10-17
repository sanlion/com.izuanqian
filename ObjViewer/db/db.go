package db

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
)

var (
	p *pool.Pool
)

func init() {

	p, _ = pool.NewCustom(
		"tcp",
		"127.0.0.1:6379",
		10,
		func(network, addr string) (*redis.Client, error) {
			c, err := redis.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			// set redis.db = 1
			if err := c.Cmd("select", 1).Err; err != nil {
				return nil, err
			}
			return c, nil
		})
}

func GetOptConnection() *redis.Client {
	client, err := p.Get()
	if err != nil {
		log.Panic(err)
	}
	return client
}
