package async_op

import "sync"

const (
	WorkerNum = 2048 //worker 里面的 chan 数量
)

var workerArray = [WorkerNum]*worker{}

var initWorkerLocker = &sync.Mutex{}

func Process(bindId int, asyncOp func(), continueWithOp func()) {
	if asyncOp == nil {
		return
	}

	currWorker := getCurrWorker(bindId)
	if currWorker != nil {
		currWorker.process(asyncOp, continueWithOp)
	}
}

//根据 bindId 获取一个工人
func getCurrWorker(bindId int) *worker {
	if bindId < 0 {
		bindId = -bindId
	}

	workerIdx := bindId % len(workerArray)
	currWorker := workerArray[workerIdx]

	//如果当前 worker 不为空，则直接返回
	if currWorker != nil {
		return currWorker
	}

	initWorkerLocker.Lock()
	defer initWorkerLocker.Unlock()

	//需要再次判断一下，如果不为空，则直接返回
	currWorker = workerArray[workerIdx]

	if currWorker != nil {
		return currWorker
	}

	currWorker = &worker{
		taskQ: make(chan func(), WorkerNum), //taskQ = new LinkedBlockingQueue<Function>();
	}

	//让 worker 开始工作
	go currWorker.loopExecTask()

	return currWorker
}
