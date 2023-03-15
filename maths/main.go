package main

import (
	"os"
	"time"

	"github.com/huytran2000-hcmus/learn-go-with-tests/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
