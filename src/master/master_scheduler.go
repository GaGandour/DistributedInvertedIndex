package master

import (
	"dii/customrpc"
	"log"
	"sync"
)

// Schedules map operations on remote workers. This will run until InputFilePathChan
// is closed. If there is no worker available, it'll block.
func (master *Master) schedule(proc string) []int {
	var (
		wg        sync.WaitGroup
		worker    *RemoteWorker
		operation *Operation
	)

	master.successfulOperations = 0
	master.totalOperations = 0

	for master.totalOperations < master.numIntersections {
		set1 := <-master.intersectionChan
		set2 := <-master.intersectionChan
		operation = &Operation{proc, set1, set2}

		worker = <-master.idleWorkerChan
		wg.Add(1)
		go master.runOperation(worker, operation, &wg)
		master.totalOperations++
	}

	wg.Wait()
	return <-master.intersectionChan
}

// runOperation start a single operation on a RemoteWorker and wait for it to return or fail.
func (master *Master) runOperation(remoteWorker *RemoteWorker, operation *Operation, wg *sync.WaitGroup) {
	var (
		err  error
		args *customrpc.RunArgs
	)

	args = &customrpc.RunArgs{Set1: operation.set1, Set2: operation.set2}

	reply := customrpc.IntersectReply{}
	err = remoteWorker.callRemoteWorker(operation.proc, args, &reply)

	if err != nil {
		log.Printf("Operation Failed. Error: %v\n", err)
		wg.Done()
	} else {
		wg.Done()
		master.idleWorkerChan <- remoteWorker
		master.successfulOperations++
		master.intersectionChan <- reply.Result
	}
}
