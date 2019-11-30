package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestGetNTPTime1(t *testing.T) {
	testTime := time.Date(2019, time.November, 30, 23, 0, 0, 0, time.UTC)
	mockNtpTime := func(host string) (time.Time, error) {
		return testTime, nil
	}

	ntpTime := getNTPTime(mockNtpTime)
	if ntpTime != testTime {
		t.Fatalf("function GetNTPTime() returned wrong value!")
	}
}

func TestGetNTPTime2(t *testing.T) {
	if os.Getenv("TEST_GET_NTP_TIME_FAIL") == "1" {
		mockNtpTime := func(host string) (time.Time, error) {
			return time.Now(), errors.New("Some NTP error")
		}
		getNTPTime(mockNtpTime)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestGetNTPTime2")
	cmd.Env = append(os.Environ(), "TEST_GET_NTP_TIME_FAIL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("function GetNTPTime() doesn't fail with exit code 1!")
}
