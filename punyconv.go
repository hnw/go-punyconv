package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/miekg/dns/idn"
	"os"
	"regexp"
)

func main() {
	var fp *os.File
	var err error

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [filename]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	re, _ := regexp.Compile(`^xn--`)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		name := scanner.Text()
		if re.MatchString(name) {
			name = idn.FromPunycode(name)
		} else {
			name = idn.ToPunycode(name)
		}
		fmt.Println(name)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
