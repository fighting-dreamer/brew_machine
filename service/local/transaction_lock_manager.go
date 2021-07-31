package service

import (
	"strings"
	"sync"

	"nipun.io/brew_machine/logger"
)


type CounterMutex struct {
	Mutex   *sync.Mutex
	Counter int
}

func NewCounterMutex() *CounterMutex {
	return &CounterMutex{
		Mutex:   &sync.Mutex{},
		Counter: 0,
	}
}

type TransactionLockManager struct {
	Keeper      sync.Map//map[string]*sync.Mutex
	KeeperState sync.Map//map[string]string
}

const (
	Locked = "LOCKED"
)

//var mutex sync.Mutex // Global Mutex, all operation of locking and un-locking happen in a serial manner using this.

func (tlm *TransactionLockManager) AcquireLock(entities []string) {
	logger.Logger.Debug().Msgf("Trying to acquire operation lock for %+v", entities)
	//mutex.Lock()
	entitiesString := strings.Join(entities, "-")
	value, ok := tlm.Keeper.Load(entitiesString)
	if !ok {
		tlm.Keeper.Store(entitiesString, NewCounterMutex())// = &sync.Mutex{}
		logger.Logger.Debug().Msgf("Created mutex for entity : %s", entitiesString)
	}
	logger.Logger.Debug().Msgf("Trying to acquire lock for %+v", entities)
	value, ok = tlm.Keeper.Load(entitiesString)
	if ok {
		value.(*CounterMutex).Mutex.Lock()
	}
	//tlm.Keeper[entitiesString].Lock() // who ever want to acquire a lock on something already locked, will have to wait.
	logger.Logger.Debug().Msgf("Acquired lock on mutex for entity : %s", entitiesString)
	//tlm.KeeperState[entitiesString] = Locked
	tlm.KeeperState.Store(entitiesString, Locked)
	//mutex.Unlock()
}

func (tlm *TransactionLockManager) ReleaseLock(entities []string) {
	//mutex.Lock() // to deal with classic deadlock scenario
	entitiesString := strings.Join(entities, "-")
	stateValue, stateOk := tlm.KeeperState.Load(entitiesString)
	if !stateOk &&  stateValue == Locked {
		logger.Logger.Debug().Msgf("released lock on mutex for entity : %s", entitiesString)
		value, ok := tlm.Keeper.Load(entitiesString)
		if ok {
			value.(*CounterMutex).Mutex.Unlock()
			value.(*CounterMutex).Counter++
			tlm.KeeperState.Store(entitiesString, "")
		}
		//tlm.Keeper[entitiesString].Unlock()
		//tlm.KeeperState[entitiesString] = ""
	}
	//mutex.Unlock()
}

func handler(key, value interface{}) bool {
	logger.Logger.Debug().Msgf("Name :%s %s\n", key, value)
	return true
}

func (tlm *TransactionLockManager) GetCurrentLockState() {
	//mutex.Lock()
	tlm.KeeperState.Range(handler)
	//for k, v := range tlm.KeeperState.Range() {
	//	logger.Logger.Debug().Msgf("Key : %s Value : %s", k, v)
	//}
	//mutex.Unlock()
}

