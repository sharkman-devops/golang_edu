package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

type ntpTime func(host string) (time.Time, error)

func getNTPTime(ntpTimeFunc ntpTime) time.Time {
	ntpTimeValue, err := ntpTimeFunc("ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return ntpTimeValue
}

func main() {
	fmt.Println(getNTPTime(ntp.Time))
}
