package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os env: %+v\n", os.Getenv("APP_ENV"))
}
