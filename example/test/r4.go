package main

import (
	"fmt"
	"github.com/hasuburero/japanese/japanese"
)

func main() {
	buf, err := japanese.IndexConbination([]int{}, 1, 4)
	if err != nil {
		fmt.Println(err)
		fmt.Println("japanese.IndexConbination error")
		return
	}
	fmt.Println(buf)
	return
}
