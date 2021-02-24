package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		from  = flag.String("from", "", "input directory folder.")
		to    = flag.String("to", "", "output file folder.")
		mode  = flag.String("mode", "search", "symmetrix or date")
		force = flag.Bool("force", false, "force overwrite copy")
	)

	flag.Parse()

	if *from == "" {
		log.Fatalln("Please input the target directory name.")
	}

	if *to == "" {
		log.Fatalln("Please input the target directory name.")
	}

	if *mode == "" {
		log.Fatalln("Please output the target file path.")
	}

	fmt.Println(force)

}
