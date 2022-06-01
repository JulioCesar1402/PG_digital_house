package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("momento do error: %v", time.Now())
	fmt.Println("error:", err)
}