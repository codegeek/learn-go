package main

import (
	"os"
	"time"

	"github.com/CodeGeek/learn-go/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
