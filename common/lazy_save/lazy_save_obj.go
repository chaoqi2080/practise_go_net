package lazy_save

type LazySaveObj interface {
	GetLsoId() string
	GetLastUpdateTime() int64
	SetLastUpdateTime(val int64)
}
