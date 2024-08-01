package main

import (
	"fmt"

	"github.com/IvanDobriy/Module1/pkg/user"
)

func main() {
	module := user.MyModule{}
	fmt.Println("Module Version:", module.Version())
}
