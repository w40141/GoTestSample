package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ゴルーチンを3つ起動
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	// メインゴルーチンを一定時間スリープさせる
	defer fmt.Println("Main goroutine has finished.")
	wg.Wait()
}

func worker(wg *sync.WaitGroup, id int) {
	fmt.Printf("Worker %d started\n", id)
	// 一つだけゴルーチンをスリープさせる
	defer wg.Done()
	if id == 1 {
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("Worker %d finished\n", id)
}
