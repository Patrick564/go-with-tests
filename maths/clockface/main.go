package main

import (
	"os"
	"time"

	clockface "github.com/Patrick564/go-with-tests/maths"
)

func main() {
	t := time.Now()

	clockface.SVGWriter(os.Stdout, t)
}
