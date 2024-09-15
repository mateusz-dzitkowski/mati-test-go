package main

import (
	"redis/resp"
	"sync"
)

var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	"HSET": hset,
	"HGET": hget,
}

func ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: resp.STRING, Str: "PONG"}
	}
	return resp.Value{Typ: resp.STRING, Str: args[0].Bulk}
}

var SETs = map[string]string{}
var SETsMutex = &sync.RWMutex{}

func set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: resp.ERROR, Str: "invalid arguments"}
	}
	key := args[0].Bulk
	value := args[1].Bulk

	SETsMutex.Lock()
	SETs[key] = value
	SETsMutex.Unlock()

	return resp.Value{Typ: resp.STRING, Str: "OK"}
}

func get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: resp.ERROR, Str: "invalid arguments"}
	}

	key := args[0].Bulk
	SETsMutex.RLock()
	val, ok := SETs[key]
	SETsMutex.RUnlock()

	if !ok {
		return resp.Value{Typ: resp.NULL}
	}
	return resp.Value{Typ: resp.STRING, Str: val}
}

var HSETs = map[string]map[string]string{}
var HSETsMutex = &sync.RWMutex{}

func hset(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: resp.ERROR, Str: "invalid arguments"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMutex.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMutex.Unlock()

	return resp.Value{Typ: resp.STRING, Str: "OK"}
}

func hget(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: resp.ERROR, Str: "invalid arguments"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMutex.RLock()
	value, ok := HSETs[hash][key]
	HSETsMutex.RUnlock()

	if !ok {
		return resp.Value{Typ: resp.NULL}
	}
	return resp.Value{Typ: resp.STRING, Str: value}
}
