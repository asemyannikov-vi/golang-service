package main

import (
	"fmt"

	tee "github.com/asemyannikov-vi/tee"
)

func main() {
	fmt.Println("Message From tee:")
	fmt.Println(tee.GetCommonData())
}
