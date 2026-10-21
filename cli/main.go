// Command signalhunt is a command-line Bluetooth Low Energy (BLE) proximity
// finder that locates nearby devices using real-time signal strength.
package main

import (
	"fmt"
	"os"

	"github.com/Nagaraj-Sunagar-Labs/signalhunt/scan"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "signalhunt: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fmt.Println("signalhunt: Starting BLE proximity finder...")
	scan.ScanForDevices()
	return nil
}
