package main

import (
	"where-am-i-server/Routes"
)

func main() {
    router := Routes.SetupRouter()
	router.Run("localhost:8080")
}
