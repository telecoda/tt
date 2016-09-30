package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var nanoTime int64
var unixTime int64
var textTime string

func init() {

	flag.Int64Var(&nanoTime, "n", 0, "time in nanoseconds")
	flag.Int64Var(&unixTime, "u", 0, "time in unix")
	flag.StringVar(&textTime, "s", "", fmt.Sprintf("time in format %s", time.RFC3339))
}

func getTime() (time.Time, error) {
	if nanoTime != 0 {
		return time.Unix(0, nanoTime), nil
	}
	if unixTime != 0 {
		return time.Unix(unixTime, 0), nil
	}
	if textTime != "" {
		// lets Try and parse this string..
		return time.Parse(time.RFC3339, textTime)
	}

	return time.Now(), nil
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("\n  Simple time conversion tool.\n")
		fmt.Printf("  Provide a time input in a variety of formats\n")
		fmt.Printf("  If no time provided, defaults to current time\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	timestamp, err := getTime()

	if err != nil {
		fmt.Printf("Error parsing parameters: %s\n", err)
		return
	}

	fmt.Printf("Format      :                     Value\n")
	fmt.Printf("=======================================\n")
	fmt.Printf("Nano        : %25d\n", timestamp.UnixNano())
	fmt.Printf("Unix        : %25d\n", timestamp.Unix())
	fmt.Printf("RFC3339     : %s\n", timestamp.Format(time.RFC3339))
	fmt.Printf("RFC3339Nano : %s\n", timestamp.Format(time.RFC3339Nano))

}
