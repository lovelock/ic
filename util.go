package main

import (
	"math/rand"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomPortForLocalForward() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(10000)
}
