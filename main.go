package main

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func main() {

	AuthenticationService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := AuthenticationService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
	}

	ErrorStartService := AuthenticationService.Start()
	if ErrorStartService != nil {

		fmt.Println(ErrorStartService)
	}
}
