// Copyright © 2018 Martin Strobel

package main

import (
	"github.com/gobuffalo/envy"
	"github.com/marstr/gophercon2018-cloudnative/exercises/service_bus/sudoku_subscriber/cmd"
)

func main() {
	envy.Load(".env")
	cmd.Execute()
}
