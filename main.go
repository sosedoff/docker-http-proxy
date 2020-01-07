package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    target := os.Getenv("TARGET_URL")
    if target == "" {
        log.Fatal("TARGET_URL is not provided")
    }

    targetURL, err := url.Parse(target)
    if err != nil {
        log.Fatal(err)
    }

    proxy := httputil.NewSingleHostReverseProxy(targetURL)

    if err := http.ListenAndServe("0.0.0.0:"+port, proxy); err != nil {
        log.Fatal(err)
    }
}

