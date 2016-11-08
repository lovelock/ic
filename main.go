package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//Ssh("hadoop", "1", "192.168.159.3", "22")
	cmd := exec.Command("ls", "-l")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
