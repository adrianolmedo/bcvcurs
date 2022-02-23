package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cfg *Config

func main() {
	// Flag or Env? Why not both for port?
	defaultPort := 80
	envPort, ok := os.LookupEnv("PORT")
	if ok {
		if p, err := strconv.Atoi(envPort); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		} else {
			defaultPort = p
		}
	}

	var (
		addr = flag.String("addr", getNetworkIP(), "Internal address of the container")
		port = flag.Int("port", defaultPort, "Internal port of the container")
	)
	flag.Parse()

	cfg = &Config{
		Addr: *addr,
		Port: *port,
		Logger: NewDebug(func(d *Debug) {
			d.timefmt = "2006-01-02T15:15:04:05"
		}),
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("/", mGET(router))

	s := &http.Server{
		Addr:           cfg.Addr + ":" + strconv.Itoa(cfg.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	cfg.Logger.Log("level", "info", "addr", cfg.Addr+":"+strconv.Itoa(cfg.Port), "msg", "listening")
	return s.ListenAndServe()
}

// getNetworkIP return local network IP. If you are not connected to IPv4 it will return 127.0.0.1.
func getNetworkIP() string {
	ip := "127.0.0.1"

	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
			ip := networkIP.IP.String()
			return ip
		}
	}

	return ip
}
