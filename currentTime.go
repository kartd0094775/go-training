package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-01-02 15:04:05"
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format(format))
}
