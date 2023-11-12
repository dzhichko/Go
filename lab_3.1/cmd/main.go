package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	calc "lab_3.1/calculations"
)

// Calculate n!
func main() {

	var n int64
	var err error

	switch len(os.Args) {
	case 2:
		n, err = strconv.ParseInt(os.Args[1], 10, 64)
	case 3:
		n, err = strconv.ParseInt(os.Args[2], 10, 64)
	default:
		log.Fatal("Invalid amount of Args")
	}
	if err != nil {
		log.Fatal("Invalid input")
	}
	f := flag.Bool("log", false, "Enable logging")
	flag.Parse()

	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	log.Info(calc.Calculate(n, *f))

}
