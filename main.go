package main

import (
    "echo_api/routes"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    e := routes.Init()
    e.Logger.Fatal(e.Start(":" + port))
}
