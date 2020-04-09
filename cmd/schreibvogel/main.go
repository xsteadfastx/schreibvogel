package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/xsteadfastx/schreibvogel"
)

func banner() {
	fmt.Println("                                                                                        ")
	fmt.Println("=====================================================================================   ")
	fmt.Println("                                                                                        ")
	fmt.Println("   BRINGING INDIEWEB TO XSTEADFASTX.ORG                                                 ")
	fmt.Println("  / ___/  /  ]|  T  T|    \\   /  _]l    j|    \\ |  T  | /   \\  /    T  /  _]| T      ")
	fmt.Println(" (   \\_  /  / |  l  ||  D  ) /  [_  |  T |  o  )|  |  |Y     YY   __j /  [_ | |        ")
	fmt.Println("  \\__  T/  /  |  _  ||    / Y    _] |  | |     T|  |  ||  O  ||  T  |Y    _]| l___     ")
	fmt.Println("  /  \\ /   \\_ |  |  ||    \\ |   [_  |  | |  O  |l  :  !|     ||  l_ ||   [_ |     T  ")
	fmt.Println("  \\    \\     ||  |  ||  .  Y|     T j  l |     | \\   / l     !|     ||     T|     |  ")
	fmt.Println("   \\___j\\____jl__j__jl__j\\_jl_____j|____jl_____j  \\_/   \\___/ l___,_jl_____jl_____j")
	fmt.Println("                                                           THAT B*TCH C*ROL B*SKIN      ")
	fmt.Println("=====================================================================================   ")
	fmt.Println("                                                                                        ")
}

func main() {
	banner()

	args := os.Args

	if len(args) != 3 { //nolint:gomnd
		log.Fatal("Usage: schreibvogel syndicate CONFIG")
	}

	switch {
	case args[1] == "syndicate":
		schreibvogel.Syndicate()
	default:
		log.Fatal("nothing to do")
	}
}
