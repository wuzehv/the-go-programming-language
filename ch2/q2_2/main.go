package main

import (
	"flag"
	"fmt"
	"github.com/wuzehv/the_go_programming_language/ch2/q2_1"
	"os"
	"strconv"
)

func main() {
	ck := flag.Bool("ck", false, "Celsius To Kelvin")
	kc := flag.Bool("kc", false, "Kelvin To Celsius")
	cf := flag.Bool("cf", false, "Celsius To Fahrenheit")
	fc := flag.Bool("fc", false, "Fahrenheit To Celsius")

	flag.Parse()

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	t, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	switch {
	case *ck:
		fmt.Println(q2_1.CToK(q2_1.Celsius(t)))
	case *kc:
		fmt.Println(q2_1.KToC(q2_1.Kelvin(t)))
	case *cf:
		fmt.Println(q2_1.CToF(q2_1.Celsius(t)))
	case *fc:
		fmt.Println(q2_1.FToC(q2_1.Fahrenheit(t)))
	default:
		flag.Usage()
		os.Exit(1)
	}
}
