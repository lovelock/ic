package main

import (
	"math/rand"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomPortForLocalForward() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := strconv.FormatInt(r.Int63n(10000), 10)
	return s
}
