package main

import (
	"belajar-gin/routes"
)

func main() {
	PORT := ":3001"

	routes.StartServer().Run(PORT)
}
