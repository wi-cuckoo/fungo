package main

import (
	"fmt"
	"os"
)

const banner = `
    _______           ________     
   / ____(_)_______  / ____/ /_  __
  / /_  / / ___/ _ \/ /_  / / / / /
 / __/ / / /  /  __/ __/ / / /_/ / 
/_/   /_/_/   \___/_/   /_/\__, /  
                          /____/  hello, my lord! 
`

func init() {
	fmt.Fprint(os.Stdout, "\033c")
	// filepath.Walk("~/.poems", nil)
}

func main() {
	fmt.Fprintln(os.Stdout, banner)
}
