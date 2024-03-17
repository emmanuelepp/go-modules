package main

import (
	"fmt"
	"toolset"
)

func main() {

	var tools toolset.Tools

	s := tools.RandomPassword(16)

	fmt.Println("Random Password: ", s)

}
