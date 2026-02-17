// NOTICE

// Project Name: Cloaq
// Copyright © 2026 Neil Talap and/or its designated Affiliates.

// This software is licensed under the Dragonfly Public License (DPL) 1.0.

// All rights reserved. The names "Neil Talap" and any associated logos or branding
// are trademarks of the Licensor and may not be used without express written permission,
// except as provided in Section 7 of the License.

// For commercial licensing inquiries or permissions beyond the scope of this
// license, please create an issue in github.

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloaq <command>")
		return
	}

	switch os.Args[1] {
	case "run":
		runCommand()
	case "settings":
		settingsCommand()
	case "help":
		helpCommand()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func runCommand() {
	fmt.Println("Running Cloaq")

}

func helpCommand() {
	fmt.Println("help text")
}

func settingsCommand() {
	fmt.Println("settings text")
}
