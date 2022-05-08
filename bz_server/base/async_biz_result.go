package base

import (
	"practise_go_net/common/main_thread"
	"sync/atomic"
)

type AsyncBizResult struct {
	//已经返回对象
	returnedObj interface{}
	//完成回调函数
	completeFunc func()

	//是否已经设置对象 0:没有设置，1:已经设置
	hasReturnedObj int32
	//是否已经设置回调函数 0:没有设置，1:已经设置
	hasCompleteFunc int32
	//是否已经调用回调 0:没有设置，1:已经设置
	completeFuncHasAlreadyBeenCalled int32
}

func (bizResult *AsyncBizResult) GetReturnedObj() interface{} {
	return bizResult.returnedObj
}

func (bizResult *AsyncBizResult) SetReturnedObj(val interface{}) {
	if atomic.CompareAndSwapInt32(&bizResult.hasReturnedObj, 0, 1) {
		bizResult.returnedObj = val
		//看是否已经设置了回调函数，有则执行回调
		bizResult.doComplete()
	}
}

func (bizResult *AsyncBizResult) OnComplete(fn func()) {
	if atomic.CompareAndSwapInt32(&bizResult.hasCompleteFunc, 0, 1) {
		bizResult.completeFunc = fn
		//看是否已经设置了返回对象，有就直接调用
		if bizResult.hasReturnedObj == 1 {
			bizResult.doComplete()
		}
	}
}

func (bizResult *AsyncBizResult) doComplete() {
	//回调函数没设置，直接返回
	if bizResult.completeFunc == nil {
		return
	}

	if atomic.CompareAndSwapInt32(&bizResult.completeFuncHasAlreadyBeenCalled, 0, 1) {
		main_thread.Process(bizResult.completeFunc)
	}
}
