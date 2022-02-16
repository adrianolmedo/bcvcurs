package main

import (
	"bytes"
	"flag"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cfg *Config

func main() {
	var (
		addr = flag.String("addr", getNetworkIP(), "Internal address of the container")
		port = flag.Int("port", 80, "Internal port of the container")
	)
	flag.Parse()

	cfg = &Config{
		Addr:   *addr,
		Port:   *port,
		Logger: NewDebug(&bytes.Buffer{}),
	}

	if err := run(); err != nil {
		cfg.Logger.Log("level", "error", "msg", err.Error())
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("/", mGET(home()))
	http.HandleFunc("/v1/", mGET(home()))
	http.HandleFunc("/v1/dollar/", mGET(dollar(cfg.Logger)))

	s := &http.Server{
		Addr:           cfg.Addr + ":" + strconv.Itoa(cfg.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	cfg.Logger.Log("level", "info", "addr", cfg.Addr, "port", cfg.Port, "msg", "listening")
	return s.ListenAndServe()
}

// getNetworkIP return local network IP. If you are not connected to IPv4 it will return empty string.
func getNetworkIP() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
			ip := networkIP.IP.String()
			return ip
		}
	}
	return ""
}
