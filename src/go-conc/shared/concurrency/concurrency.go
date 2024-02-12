package concurrency

import "sync"

type CountAndWait struct {
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
	counter   int
}

func (caw *CountAndWait) RoutineStart() {
	caw.waitGroup.Add(1)
	caw.mutex.Lock()
	caw.counter++
	caw.mutex.Unlock()
}

func (caw *CountAndWait) RoutineDone() {
	caw.waitGroup.Done()
}

func (caw *CountAndWait) WaitAllRoutinesDone() {
	caw.waitGroup.Wait()
}

func (caw *CountAndWait) Counter() int {
	return caw.counter
}
