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

    router := handlers.Router(version.BuildTime, version.Commit, version.Release)
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

    srv := &http.Server{
        Addr:    ":" + port,
        Handler: r,
    }

    go func() {
        log.Fatal(srv.ListenAndServe())
    }()

    log.Print("The service is ready to listen and serve on port: "+port)
    log.Fatal(http.ListenAndServe(":"+port, router))

    killSignal := <-interrupt
    switch killSignal {
      case os.Interrupt:
        log.Print("Got SIGINT...")
      case syscall.SIGTERM:
        log.Print("Got SIGTERM...")
    }
    log.Print("The service is shutting down...")
    srv.Shutdown(context.Background())
    log.Print("Done")
}
