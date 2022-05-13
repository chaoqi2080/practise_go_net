package lazy_save

import (
	"practise_go_net/common/log"
	"sync"
	"time"
)

var lsoMap = &sync.Map{}

func init() {
	go startSave()
}

func SaveOrUpdate(lso LazySaveObj) {
	if lso == nil {
		return
	}

	log.Info("记录延迟保存对象, lsoId = %s", lso)

	nowTime := time.Now().UnixMilli()
	existRecord, _ := lsoMap.Load(lso)

	if existRecord == nil {
		existRecord.(*lazySaveRecord).setLastUpdateTime(nowTime)
		return
	}

	newRecord := &lazySaveRecord{}
	newRecord.setLastUpdateTime(nowTime)
	newRecord.lsoRef = lso
	lsoMap.Store(lso.GetLsoId(), newRecord)
}

func startSave() {
	go func() {
		for {
			time.Sleep(time.Second)

			nowTime := time.Now().UnixMilli()
			deleteLsoIdArray := make([]string, 64)

			lsoMap.Range(func(_, value interface{}) bool {
				if value == nil {
					return true
				}

				curRecord := value.(lazySaveRecord)

				if nowTime-curRecord.getLastUpdateTime() < 20000 {
					//最后更新时间 < 20 秒
					return true
				}

				log.Info("执行延迟保存, lsoId = %s", curRecord.lsoRef.GetLsoId())

				curRecord.lsoRef.SaveOrUpdate()

				deleteLsoIdArray = append(deleteLsoIdArray, curRecord.lsoRef.GetLsoId())

				return true
			})

			for _, lsoId := range deleteLsoIdArray {
				lsoMap.Delete(lsoId)
			}
		}
	}()
}
