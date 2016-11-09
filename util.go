package main

import ()

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
