package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Copy func copies file from src `from` to dst `to`
// Args:
// `limit` - limit in bytes of `from` file. Default: inf bytes (whole file)
// `offset` - offset of `from` file in bytes. Default: 0 bytes
func Copy(from string, to string, limit int, offset int) error {
	if limit < 1 {
		return errors.New("limit must be greater than 0")
	}
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	stat, err := src.Stat()
	if err != nil {
		return err
	}
	if ! stat.Mode().IsRegular() {
		return errors.New("only regular files allowed")
	}
	if stat.Size() < int64(offset+limit) {
		return errors.New("offset + limit must be lower than file length")
	}

	newOffset, err := src.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	if newOffset != int64(offset) {
		return errors.New("failed to set offset")
	}

	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dst.Close()

	written, err := io.CopyN(dst, src, int64(limit))
	if err != nil {
		return err
	}

	fmt.Printf("%v bytes copied", written)
	return nil
}

func main1() {
	err := Copy("/tmp/src", "/tmp/dst", 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
