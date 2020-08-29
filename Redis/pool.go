// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"log"
	"time"

	"github.com/FZambia/go-sentinel"
	"github.com/gomodule/redigo/redis"
)

var RedisConnPool *redis.Pool

var Stnl sentinel.Sentinel

func InitPool() {
	Stnl = sentinel.Sentinel{
		Addrs:      []string{"127.0.0.1:26379", ":26380", ":26381"},
		MasterName: "mymaster",
		Dial: func(addr string) (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", ":26380", redis.DialConnectTimeout(time.Second*3))
			if err != nil {
				log.Fatalln("Pool:",err)
			}
			return conn, nil
		},
		Pool: nil,
	}

	// Dial：新建连接的工厂函数。
	// TestOnBorrow：连接的健康检测函数。
	// MaxIdle：最大空闲连接数。
	// MaxActive：最大活跃连接数。
	// IdleTimeout：空闲连接的超时时间
	// Wait：如果连接池达到了最大的活跃连接数，Wait用以指示是否需要继续等待。
	// active：活跃连接数量。
	// closed：连接池是否关闭。
	// idle：维护空闲连接的集合，idleList的实现与链表类似，不多介绍。
	RedisConnPool = &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 300 * time.Second,
		MaxActive:   500,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := Stnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: CheckRedisRole,
	}
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}
}
