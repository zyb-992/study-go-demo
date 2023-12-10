package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

}

type TaskInfo struct {
	count int
	wg    *sync.WaitGroup
}

var taskInfo = new(TaskInfo)

func Add() {
	taskInfo.count++
	taskInfo.wg.Add(1)
}

func Shutdown() {
	taskInfo.wg.Wait()
}

func Completed() {
	taskInfo.wg.Done()
}

func gracefulExit() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	s := <-ch
	log.Println("catch signal: ", s)

	// ... 处理逻辑
	Shutdown()
}
