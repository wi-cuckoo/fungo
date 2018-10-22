// thank to http://patorjk.com/software/taag

package main

import (
	"fmt"
)

const banner = `
    _______           ________     
   / ____(_)_______  / ____/ /_  __
  / /_  / / ___/ _ \/ /_  / / / / /
 / __/ / / /  /  __/ __/ / / /_/ / 
/_/   /_/_/   \___/_/   /_/\__, /  
                          /____/  hello, my lord! 
`

func banner() {
	fmt.Println("\033[H\033[2J")
	fmt.Println(banner)
}
