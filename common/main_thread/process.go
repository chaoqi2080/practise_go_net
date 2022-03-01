package main_thread

import "sync"

const mainQSize = 2048

var mainQ = make(chan func(), mainQSize)

var started = false
var startedLocker = &sync.Mutex{}

func Process(task func()) {
	if task == nil {
		return
	}

	//启动处理协程
	if !started {
		startedLocker.Lock()
		defer startedLocker.Unlock()

		if !started {
			started = true
			mainQ <- task

			go execute()
		}
	}
}

func execute() {
	for {
		task := <-mainQ

		if task != nil {
			task()
		}
	}
}
