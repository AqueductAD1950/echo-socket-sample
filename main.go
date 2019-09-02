package main

import (
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	socketFile := "/tmp/goecho.sock"
	os.Remove(socketFile)

	l, err := net.Listen("unix", socketFile)
	if err != nil {
		e.Logger.Fatal(err)
	}

	err = os.Chmod(socketFile, 0777)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Listener = l
	e.Logger.Fatal(e.Start(""))
}
