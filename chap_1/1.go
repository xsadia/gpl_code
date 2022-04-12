package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func idk() {
	s, sep := "", ""

	start := time.Now()
	for i, v := range os.Args {
		s += sep + fmt.Sprintf("%d, %s", i, v)
		sep = "\n"
	}
	fb := time.Since(start).Seconds()

	start = time.Now()

	s = strings.Join(os.Args, " ")

	sb := time.Since(start).Seconds()

	fmt.Printf("For loop: %.9f\nJoin: %.9f\n", fb, sb)

	if fb < sb {
		fmt.Println("For faster")
	} else {
		fmt.Println("Join faster")
	}

	fmt.Println(s)
}
