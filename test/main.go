package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/google/go-github/github"
)

func main() {
	var x *github.Client
	if x == nil {
		fmt.Println("do nothing")
	}
}

func test1() {
	var out [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
}

func test2() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		for j := 0; j < 1; j++ {
			wg.Add(1)
			go func() {
				fmt.Printf("%d ", i)
				wg.Done()
			}()
		}
	}
	wg.Wait()
}

func test3() {
	vs := []int{1, 2, 3, 4}
	for _, v := range vs {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(1 * time.Second)
}
