package main

import "fmt"

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildTag     = "N/A"
)

func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build tag: %s\n", buildTag)
}
