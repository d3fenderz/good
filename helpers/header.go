package helpers

import (
	"fmt"

	"github.com/jmau111/good/libgood"
)

func Header() {
	fmt.Print(" ________  ________  ________  ________\n")
	fmt.Print("|\\   ____\\|\\   __  \\|\\   __  \\|\\   ___ \\\n")
	fmt.Print("\\ \\  \\___|\\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\_|\\ \\\n")
	fmt.Print(" \\ \\  \\  __\\ \\  \\\\\\  \\ \\  \\\\\\  \\ \\  \\ \\\\ \\\n")
	fmt.Print("  \\ \\  \\|\\  \\ \\  \\\\\\  \\ \\  \\\\\\  \\ \\  \\_\\\\ \\\n")
	fmt.Print("   \\ \\_______\\ \\_______\\ \\_______\\ \\_______\\\n")
	fmt.Print("    \\|_______|\\|_______|\\|_______|\\|_______|\n")
	fmt.Printf("\nv%s by jmau111\n", libgood.VERSION)
	fmt.Print("-----------------\n\n")
}
