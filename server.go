package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	PORT string = "6390"
	IP   string = "127.0.0.1"
)

type CachedPair struct {
	value          interface{}
	ExpirationTime int64
}

type LocalCache struct {
	Pairs map[string]CachedPair
	Mute  sync.Mutex
	Wg    sync.WaitGroup
	Quit  chan bool
}

type LoCache interface {
	StartCleanup(interval time.Duration)
	StopCleanup()
	PING() string
	SET(inp string) string
	GET(inp string) string
	SETNX(inp string) string
	//MSET()
	//MGET()
	//DEL()
	//INCR()
	//DECR()
	//EXPIRE()
	//TTL()
	//SETEX()
}

func getAddress() string {
	return IP + ":" + PORT
}

func newLocalCache() *LocalCache {
	lcache := &LocalCache{
		Pairs: make(map[string]CachedPair),
		Quit:  make(chan bool),
	}
	lcache.Wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lcache.Wg.Done()
		lcache.StartCleanup(cleanupInterval)
	}(500 * time.Millisecond)

	return lcache
}

func (lc *LocalCache) StartCleanup(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.Quit:
			return
		case <-t.C:
			lc.Mute.Lock()
			for k, v := range lc.Pairs {
				if v.ExpirationTime <= time.Now().Unix() && v.ExpirationTime != 0 {
					delete(lc.Pairs, k)
				}
			}
			lc.Mute.Unlock()
		}
	}
}

func (lc *LocalCache) StopCleanup() {
	close(lc.Quit)
	lc.Wg.Wait()
}

func (lc *LocalCache) PING() string {
	return "PONG"
}

func (lc *LocalCache) SET(inp string) string {
	res := "Key Added"
	inpSplitted := strings.Split(inp, " ")
	key := inpSplitted[1]
	var val interface{}
	v, err := strconv.Atoi(inpSplitted[2])
	if err == nil {
		val = v
	} else {
		tmp := inpSplitted[2]
		if tmp[0] == 34 && tmp[len(tmp)-1] == 34 {
			val = tmp[1 : len(tmp)-1]
		} else if tmp[0] == 34 || tmp[len(tmp)-1] == 34 {
			return "Invalid Value"
		} else {
			val = tmp
		}
	}
	if pop, ok := lc.Pairs[key]; !ok {
		lc.Pairs[key] = CachedPair{value: val, ExpirationTime: 0}
	} else {
		ex := pop.ExpirationTime
		lc.Pairs[key] = CachedPair{value: val, ExpirationTime: ex}
		res = "Key Updated"
	}
	return res
}

func (lc *LocalCache) GET(inp string) string {
	inpSplitted := strings.Split(inp, " ")
	key := inpSplitted[1]
	if pop, ok := lc.Pairs[key]; ok {
		return pop.value.(string)
	}
	return "Not Found"
}

func (lc *LocalCache) SETNX(inp string) string {
	res := "Key Added"
	inpSplitted := strings.Split(inp, " ")
	key := inpSplitted[1]
	var val interface{}
	v, err := strconv.Atoi(inpSplitted[2])
	if err == nil {
		val = v
	} else {
		tmp := inpSplitted[2]
		if tmp[0] == 34 && tmp[len(tmp)-1] == 34 {
			val = tmp[1 : len(tmp)-1]
		} else if tmp[0] == 34 || tmp[len(tmp)-1] == 34 {
			return "Invalid Value"
		} else {
			val = tmp
		}
	}
	if _, ok := lc.Pairs[key]; !ok {
		lc.Pairs[key] = CachedPair{value: val, ExpirationTime: 0}
	} else {
		res = "Existing Key"
	}
	return res
}

func main() {
	listener, err := net.Listen("tcp", getAddress())
	if err != nil {
		fmt.Printf("Failed to bind to port %s", PORT)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on", getAddress())

	lcache := newLocalCache()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go HandleConnection(conn, lcache)
	}
}
