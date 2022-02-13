package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		addr = flag.String("addr", getNetworkIP(), "Internal address of the container")
		port = flag.Int("port", 80, "Internal port of the container")
	)
	flag.Parse()

	cfg := Config{
		Addr:   *addr,
		Port:   *port,
		Logger: NewLogger(),
	}

	if err := run(cfg); err != nil {
		cfg.Logger.Logf("%s", err)
		os.Exit(1)
	}
}

func run(cfg Config) error {
	http.HandleFunc("/", GET(home))
	http.HandleFunc("/v1/", GET(home))
	http.HandleFunc("/v1/dollar/", GET(dollar))

	s := &http.Server{
		Addr:           cfg.Addr + ":" + strconv.Itoa(cfg.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	cfg.Logger.Logf("Listening %s:%d, press Ctrl+C to exit", cfg.Addr, cfg.Port)
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
