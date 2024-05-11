package main

import "os"

func callbackExit(cfg *config, input string) error {
	os.Exit(0)
	return nil
}
