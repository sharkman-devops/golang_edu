package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cases struct {
	src                string
	dst                string
	srcContent         []byte
	limit              int
	offset             int
	dstExpectedContent []byte
	error              bool
}

var testData = []cases{
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 0, 0, []byte("0123456789"), false},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 0, 1, []byte("123456789"), false},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 1, 0, []byte("0"), false},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), -1, 0, []byte(nil), true},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 0, -1, []byte(nil), true},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 0, 10, []byte(""), false},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 0, 9, []byte("9"), false},
	{"/tmp/src_test", "/tmp/dst_test", []byte("0123456789"), 1, 10, []byte(nil), true},
	{"/tmp/src_test", "/tmp/dst_test", []byte(""), 0, 0, []byte(""), false},
}

func TestCopy(t *testing.T) {
	for i := 0; i < len(testData); i++ {
		_ = ioutil.WriteFile(testData[i].src, testData[i].srcContent, 0644)
		err := Copy(testData[i].src, testData[i].dst, testData[i].limit, testData[i].offset)
		if testData[i].error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		dat, _ := ioutil.ReadFile(testData[i].dst)
		assert.Equal(t, testData[i].dstExpectedContent, dat)

		os.Remove(testData[i].src)
		os.Remove(testData[i].dst)
	}
}
