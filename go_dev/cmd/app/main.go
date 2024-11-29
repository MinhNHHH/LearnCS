package main

import (
	"fmt"

	"github.com/MinhNHHH/go_dev/pkg/routes"
)

const port = 8080

func main() {
	route := routes.Routes()
	route.Run(fmt.Sprintf(":%d", port))
}
