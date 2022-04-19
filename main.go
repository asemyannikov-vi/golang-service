package main

import (
	"fmt"

	tee "github.com/asemyannikov-vi/tee-new"
)

func main() {
	fmt.Println("Message From:")
	fmt.Println(tee.GetReleaseMessage())
}
