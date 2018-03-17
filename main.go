package main

import (
    "log"
    "net/http"
    "os"
    "tutor/handlers"
    "tutor/version"
)

func main() {
    log.Printf(
        "Starting the service...\ncommit: %s, build time: %s, release: %s",
        version.Commit, version.BuildTime, version.Release,
    )
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("Port is not set.")
    }
    router := handlers.Router()
    log.Print("The service is ready to listen and serve on port: "+port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
