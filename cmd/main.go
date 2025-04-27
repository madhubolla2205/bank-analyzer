package main

import (
    "github.com/gin-gonic/gin"
    "github.com/madhubolla2205/bank-analyzer/pkg/server"
)

func main() {
    r := gin.Default()
    server.SetupRoutes(r)
    r.Run(":8080")
}

