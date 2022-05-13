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

	log.Info("记录延迟保存对象, lsoId = %s", lso.GetLsoId())

	nowTime := time.Now().UnixMilli()
	existLso, _ := lsoMap.Load(lso.GetLsoId())

	if existLso == nil {
		existLso.(LazySaveObj).SetLastUpdateTime(nowTime)
		return
	}

	lso.SetLastUpdateTime(nowTime)
	lsoMap.Store(lso.GetLsoId(), lso)
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

				currLso := value.(LazySaveObj)

				if nowTime-currLso.GetLastUpdateTime() < 20000 {
					//最后更新时间 < 20 秒
					return true
				}

				log.Info("执行延迟保存, lsoId = %s", currLso.GetLsoId())

				currLso.SaveOrUpdate()

				deleteLsoIdArray = append(deleteLsoIdArray, currLso.GetLsoId())

				return true
			})

			for _, lsoId := range deleteLsoIdArray {
				lsoMap.Delete(lsoId)
			}
		}
	}()
}
