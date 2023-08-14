package main

import (
	"fmt"
	"time"
)

func main() {
	// ゴルーチンを3つ起動
	for i := 0; i < 3; i++ {
		go worker(i)
	}

	// メインゴルーチンを一定時間スリープさせる
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine has finished.")
}

func worker(id int) {
	fmt.Printf("Worker %d started\n", id)
	// 一つだけゴルーチンをスリープさせる
	if id == 1 {
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("Worker %d finished\n", id)
}
