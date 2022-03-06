package async_op

import (
	"practise_go_net/common/log"
	"practise_go_net/common/main_thread"
)

type worker struct {
	taskQ chan func()
}

func (w worker) process(asyncOp func(), continueWithOp func()) {
	if w.taskQ == nil {
		log.Error("worker taskQ is empty")
		return
	}

	w.taskQ <- func() {
		asyncOp()

		//如果继续执行函数不为空，则把这个函数丢入单协程的业务协程进行处理
		if continueWithOp != nil {
			main_thread.Process(continueWithOp)
		}
	}
}

func (w worker) loopExecTask() {
	if w.taskQ == nil {
		log.Info(
			"初始化任务执行队列",
		)

		w.taskQ = make(chan func(), 2048)
	}

	for {
		curTask := <-w.taskQ

		if curTask == nil {
			continue
		}

		//执行当前任务
		curTask()
	}
}
