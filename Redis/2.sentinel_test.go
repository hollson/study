// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/FZambia/go-sentinel"
	"github.com/gomodule/redigo/redis"
)

// 1. network必须是tcp
// 2. 可传递多个Option
func TestDial(t *testing.T) {
	// Option包括授权、Tls、数据库、超时等
	conn, err := redis.Dial("tcp", ":6379", redis.DialConnectTimeout(time.Second*3))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	result, err := conn.Do("PING")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}

// 哨兵对象成员
func TestSentinel(t *testing.T) {
	stn := sentinel.Sentinel{
		Addrs:      []string{"127.0.0.1:26379", ":26380", ":26381"},
		MasterName: "mymaster",
		Dial: func(addr string) (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", ":26380", redis.DialConnectTimeout(time.Second*3))
			if err != nil {
				log.Fatalln(err)
			}
			return conn, nil
		},
		Pool: nil,
	}

	log.Println(stn.Addrs)

	// =================================
	sm, err := stn.MasterAddr()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("主服务：", stn.MasterName, sm)

	// =================================
	sb, err := stn.Slaves()
	if err != nil {
		log.Fatalln(err)
	}
	for _, value := range sb {
		log.Println("从服务：", value.Available(), value.Addr())
	}

	// =================================
	sc, err := stn.SlaveAddrs()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("从服务：", sc)

	// =================================
	sa, err := stn.SentinelAddrs()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("其他哨兵：", sa)
}

func TestPool(t *testing.T) {
	InitPool()
	for {
		fmt.Println(time.Now().Second())
		time.Sleep(time.Second * 10)

		rd := RedisConnPool.Get()
		defer rd.Close()

		ok, err := rd.Do("SET", "hello", "你好世界")
		if err != nil {
			log.Println(err)
			// Stnl.Discover()

			continue
			// log.Fatalln("Do:", err)
		}
		log.Println(ok)

		result, err := redis.String(rd.Do("GET", "hello"))
		if err != nil {
			log.Fatalln("Get:", err)
		}
		log.Println(result)

	}

}
