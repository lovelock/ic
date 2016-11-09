package main

import (
	"fmt"
	"testing"
)

func TestGetConfigContent(t *testing.T) {
	f := GetConfigContent()
	fmt.Println(f)
}

func TestListBlocks(t *testing.T) {
	ListBlocks()
}
