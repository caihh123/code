package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	f, err := os.OpenFile("../mmap.bin", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	err = syscall.Ftruncate(int(f.Fd()), 1000)
	if err != nil {
		panic(err)
	}
	data, err := syscall.Mmap(int(f.Fd()), 0, 1<<12, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(string(data))
		for k, v := range []byte("helo another word") {
			data[k] = v
		}
		time.Sleep(time.Second * 3)
	}

	err = syscall.Munmap(data)
	if err != nil {
		panic(err)
	}
}
