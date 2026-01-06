package main

import (
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8081")
}
