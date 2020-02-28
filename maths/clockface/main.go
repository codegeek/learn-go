package main

import (
	"os"
	"time"

	"github.com/codegeek/learn-go/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
