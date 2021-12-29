package main

import (
	"letscrud/src/api/endpoints/routes"
)

func main() {
	e := routes.GetEcho()
	e.Logger.Fatal(e.Start(":3000"))
}
