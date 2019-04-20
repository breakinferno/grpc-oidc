package main

import (
	"fmt"

	"github.com/breakinferno/grpc-oidc/utils/config"
	"github.com/gin-gonic/gin"
)

func main() {
	name := config.GetString("name")
	port := config.GetString("port")

	fmt.Println("mdzz")
	r := gin.Default()

	fmt.Println(name + " start listening at http://localhost:" + port)
	fmt.Printf("==> 🚀 %s listening at %s\n", name, port)
	r.Run(":" + port)
}
