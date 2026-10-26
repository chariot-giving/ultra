package main

import (
	"os"

	"github.com/chariot-giving/ultra/cli"

	_ "github.com/chariot-giving/ultra/cli/resolvers/aws"
	_ "github.com/chariot-giving/ultra/cli/resolvers/onepassword"
	_ "github.com/chariot-giving/ultra/cli/resolvers/vault"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
