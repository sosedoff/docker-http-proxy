package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type config struct {
	port         string
	target       *url.URL
	redirectRoot string
}

func initConfig() (*config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	target := os.Getenv("TARGET_URL")
	if target == "" {
		return nil, errors.New("TARGET_URL is not provided")
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	c := config{
		port:         port,
		target:       targetURL,
		redirectRoot: os.Getenv("TARGET_REDIRECT_ROOT"),
	}

	return &c, nil
}

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(cfg.target)

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Redirect the home page to some other path
		if req.URL.Path == "/" && cfg.redirectRoot != "" {
			http.Redirect(rw, req, cfg.redirectRoot, 302)
			return
		}

		// Pass request to the backend
		proxy.ServeHTTP(rw, req)
	})

	if err := http.ListenAndServe("0.0.0.0:"+cfg.port, nil); err != nil {
		log.Fatal(err)
	}
}
