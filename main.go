package main

import "github.com/MirzaHilmi/SBD-GO/app"

func main() {
	if err := app.SetupAndRun(); err != nil {
		panic(err)
	}
}
