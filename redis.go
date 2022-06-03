package main

import (
	"encoding/json"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var currentId int

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func put(value Todo) (string, error, Todo) {
	conn := newPool().Get()
	defer conn.Close()
	currentId += 1
	value.Id = currentId
	key := "TODO" + strconv.Itoa(currentId)
	jsonVal, err := json.Marshal(value)
	reply, err := redis.String(conn.Do("SET", key, jsonVal))
	return reply, err, value
}

// func getAll() Todos {
// 	var todos Todos
// 	conn := newPool().Get()
// 	defer conn.Close()
// 	reply, err := redis.Values(conn.Do("HGETALL", "TODO"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	redis.ScanStruct(&todos)
// }

func getById() {
	//TODO
}
