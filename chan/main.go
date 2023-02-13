package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sort"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	type ChanStruct struct {
		StartTime int64
		Index     int
		Num       int
	}
	chanTest := make(chan *ChanStruct, 1000)
	var M = 1000
	var N = 100

	for i := 0; i < M; i++ {
		i := i
		go func() {
			for j := 0; j < N; j++ {
				chanTest <- &ChanStruct{
					Index:     i,
					Num:       j,
					StartTime: time.Now().UnixMilli(),
				}
			}
		}()
	}
	time.Sleep(time.Second * 2)

	var max int
	var index, count int
	m := make([]int, M)

	var totalCount int
	for {
		select {
		case data := <-chanTest:
			m[data.Index]++
			if m[data.Index] > max {
				max = m[data.Index]
			}
			if data.Num == 0 {
				totalCount++
				if totalCount == M {
					sort.Slice(m, func(i, j int) bool {
						return m[i] > m[j]
					})
					fmt.Println(m)
					return
				}
			}
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)+100))
			count++
			if count >= M*N {
				fmt.Println(index, max)
				return
			}
		}
	}
}
