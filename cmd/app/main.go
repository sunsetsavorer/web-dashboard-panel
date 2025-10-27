package main

import (
	"fmt"

	"github.com/sunsetsavorer/web-dashboard-panel/internal/app"
)

func main() {

	app := app.New()

	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
