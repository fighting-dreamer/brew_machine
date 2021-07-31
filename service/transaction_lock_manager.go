package service

type TransactionLockManager interface {
	AcquireLock([]string)
	ReleaseLock([]string)
	GetCurrentLockState()
}
