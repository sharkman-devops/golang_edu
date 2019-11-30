package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

type NtpTime func(host string) (time.Time, error)

func GetNTPTime(NtpTimeFunc NtpTime) time.Time {
	ntp_time, err := NtpTimeFunc("ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return ntp_time
}

func main() {
	fmt.Println(GetNTPTime(ntp.Time))
}
