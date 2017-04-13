package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

var nanoTime int64
var unixTime int64
var textTime string

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

	return time.Time{}, fmt.Errorf("Invalid time format")
}

func main() {
	timestamp := time.Now()
	if len(os.Args) == 2 {

		parseTime(os.Args[1])
		var err error
		timestamp, err = getTime()

		if err != nil {
			fmt.Printf("Error parsing parameters: %s\n", err)
			return
		}
	}
	fmt.Printf("Format      :                     Value\n")
	fmt.Printf("=======================================\n")
	fmt.Printf("Nano        : %25d\n", timestamp.UnixNano())
	fmt.Printf("Micro       : %25d\n", timestamp.UnixNano()/1000)
	fmt.Printf("Milli       : %25d\n", timestamp.UnixNano()/1000000)
	fmt.Printf("Unix secs   : %25d\n", timestamp.Unix())
	fmt.Printf("RFC3339     : %s\n", timestamp.Format(time.RFC3339))
	fmt.Printf("RFC3339Nano : %s\n", timestamp.Format(time.RFC3339Nano))
	fmt.Printf("Code        : time.Date(%d, %d, %d, %d, %d, %d, %d, time.UTC)\n", timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second(), timestamp.Nanosecond())
}

func parseTime(tString string) {
	tInt, err := strconv.Atoi(tString)
	if err == nil {
		numDigits := math.Floor(math.Log10(float64(tInt))) + 1

		if numDigits <= 10 {
			unixTime = int64(tInt)
			return
		}

		for ii := numDigits; ii < 19; ii++ {
			tInt = tInt * 10
		}
		nanoTime = int64(tInt)

	} else {
		textTime = tString
	}

}
