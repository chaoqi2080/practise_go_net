package async_op

var workerArray = [2048]*worker{}

func Process(bindId uint64, asyncOp func(), continueWithOp func()) {
	if asyncOp == nil {
		return
	}

	workerArray[bindId].process(asyncOp, continueWithOp)

}
