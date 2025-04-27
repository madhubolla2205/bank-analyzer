package main

import (
    "github.com/madhubolla2205/bank-analyzer/pkg/server"
)

func main() {
   r := server.SetupServer()
    r.Run(":8080")
}

