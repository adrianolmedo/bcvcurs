package main

type Config struct {
	// Internal port of the container.
	Port int

	// Internal address of the container.
	Addr string

	Logger Logger
}
