/*
File Name:  Main.go
Copyright:  2020 Kleissner Investments s.r.o.
Author:     Peter Kleissner

Web server for blog.peernet.org.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

// config defines the structure of the YAML config file
var config struct {
	// Listen specifies an array of IPs (both IPv4 and IPv6 supported) or hostnames to listen on.
	Listen []string `yaml:"Listen"`
	Port   int      `yaml:"Port"`

	// SSL configuration. The certificate file contains the certificate and may optionally include the intermediate certificate as well. Data format is PEM. The certificate key is the private key.
	UseSSL          bool   `yaml:"UseSSL"`
	CertificateFile string `yaml:"CertificateFile"`
	CertificateKey  string `yaml:"CertificateKey"`

	LogFile string `yaml:"LogFile"`

	// WebFiles is the directory holding all HTML and other files to be served by the server
	WebFiles string `yaml:"WebFiles"`
}

func init() {

	var configFile string
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	} else {
		configFile = "Server.yaml"
	}

	loadConfig(configFile)
}

// loadConfig reads the configuration file and interprets it as YAML
func loadConfig(configFile string) {
	// start handling the configuration
	cfg, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Configuration file '%s' could not be found!\nPlease ensure that your current path leads to the right file.\n", configFile)
		os.Exit(1)
	}

	err = yaml.Unmarshal(cfg, &config)
	if err != nil {
		fmt.Printf("YAML data in '%s' couldn't be read!\nPlease ensure that the format of the file is valid YAML.\n", configFile)
		os.Exit(1)
	}

	// all error logging goes to a file
	logFile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error creating log file '%s': %v\n", config.LogFile, err)
	}
	//	defer logFile.Close()	// has to remain open until program closes

	// redirect all output to the log file
	log.SetOutput(logFile)
}

// startWebServer starts a web-server with given parameters and logs the status
func startWebServer(listen string, port int, useSSL bool, certificateFile, certificateKey string, server *http.Server) {

	server.Addr = net.JoinHostPort(listen, strconv.Itoa(port))
	log.Printf("Web Server to listen on %s\n", server.Addr)

	if useSSL {
		// HTTPS
		go func() {
			if err := server.ListenAndServeTLS(certificateFile, certificateKey); err != nil {
				log.Printf("Error listening on '%s': %v\n", server.Addr, err)
			}
		}()

		// redirect HTTP -> HTTPS
		go http.ListenAndServe(net.JoinHostPort(listen, "80"), http.HandlerFunc(redirect))

	} else {
		// HTTP
		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Printf("Error listening on '%s': %v\n", server.Addr, err)
			}
		}()
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

func disableDirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			// Exception of the exception: if /index.html is present, continue
			if _, err := os.Stat(config.WebFiles + "/" + r.URL.Path + "index.html"); err != nil {
				http.NotFound(w, r)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// HeadersMiddleware sets the HSTS headers according to the input. It returns a middleware function to be used with mux.Router.Use().
func HeadersMiddleware(SetHSTS bool) func(http.Handler) http.Handler {
	return (func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set HSTS header. We include sub-domains to be secured. 1 year.
			if r.TLS != nil && SetHSTS {
				w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
			}

			// Allow caching for up to 10 minutes
			w.Header().Set("Cache-Control", "private, max-age=600")

			next.ServeHTTP(w, r)
		})
	})
}

func main() {

	fileServer := http.FileServer(http.Dir(config.WebFiles))

	// define the routes where the HTTP API will listen
	router := mux.NewRouter()
	router.Use(HeadersMiddleware(config.UseSSL))
	router.PathPrefix("/").Handler(http.StripPrefix("/", disableDirectoryListing(fileServer))).Methods("GET")

	// start the server either with SSL (HTTPS) or without (HTTP)
	for _, listen := range config.Listen {
		startWebServer(listen, config.Port, config.UseSSL, config.CertificateFile, config.CertificateKey, &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
			Handler:      router,
		})
	}

	// wait forever
	select {}
}
