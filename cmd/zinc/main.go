package main

import (
	"fmt"
	"log"
	"os"
	"path"

	. "github.com/SirMetathyst/zinc/cmd/zinc/template"
)

func main() {
	if len(os.Args) >= 2 {
		if fn, ok := Generator[os.Args[1]]; ok {
			if output, err := fn(os.Args[2:]); err == nil {
				f, err := os.Create(path.Join("./", fmt.Sprint("zinc_", output.Filename)))
				defer f.Close()
				if err != nil {
					log.Fatalf("zinc: %v", err)
				}
				_, err = f.WriteString(output.Content)
				if err != nil {
					log.Fatalf("zinc: %v", err)
				}
			} else {
				log.Fatalf("zinc: %v", err)
			}
			return
		}
	}
	fmt.Printf("Usage of zinc: zinc <%s> [Param|-h] ...\n", PrintCommands())
}
