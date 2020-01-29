package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

type barCounter struct {
	bar *pb.ProgressBar
}

func (bc *barCounter) Write(bytes []byte) (int, error) {
	bc.bar.Add(len(bytes))
	return len(bytes), nil
}

var bar *barCounter
var once sync.Once

func getBar(total int64) *barCounter {
	once.Do(func() { bar = &barCounter{pb.Start64(total)} })
	return bar
}

// Copy func copies file from src `from` to dst `to`
// Args:
// `limit` - limit in bytes of `from` file. Default: 0 bytes (whole file)
// `offset` - offset of `from` file in bytes. Default: 0 bytes
func Copy(from string, to string, limit int, offset int) error {
	if limit < 0 {
		return errors.New("limit must be positive value. 0 - whole file")
	}

	offset64 := int64(offset)
	limit64 := int64(limit)

	// src
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	stat, err := src.Stat()
	if err != nil {
		return err
	}
	if !stat.Mode().IsRegular() {
		return errors.New("only regular files allowed")
	}
	if stat.Size() < (offset64 + limit64) {
		return errors.New("offset + limit must be lower than file length")
	}
	if limit64 == 0 {
		limit64 = stat.Size() - offset64
	}

	newOffset, err := src.Seek(offset64, 0)
	if err != nil {
		return err
	}

	if newOffset != offset64 {
		return errors.New("failed to set offset")
	}

	// dst
	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dst.Close()

	// count
	fmt.Println(limit64)
	bc := getBar(limit64)
	defer bc.bar.Finish()
	srcReader := io.TeeReader(src, bc)

	// copy
	written, err := io.CopyN(dst, srcReader, limit64)
	if err != nil {
		return err
	}

	if written != limit64 {
		return fmt.Errorf("copied %v of %v bytes", written, limit64)
	}
	//fmt.Printf("%v bytes copied", written)
	return nil
}

func main() {
	err := Copy("/tmp/src", "/tmp/dst", 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
