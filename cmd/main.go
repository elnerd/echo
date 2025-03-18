package main

import "github.com/elnerd/echo/v4"

func main() {
	e := echo.New()
	e.Server.Addr = ":8080"
	e.Start("localhost:8080")

}
