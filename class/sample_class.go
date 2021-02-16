package main

import (
	"fmt"

	"github.com/bborbe/sample_go/sample_class/user"
)

func main() {
	u := user.New()
	u.Firstname = "Ben"
	u.Lastname = "Bo"
	fmt.Println(u.DisplayName())
}
