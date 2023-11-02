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

	master.failedOperationChan = make(chan *Operation, RETRY_OPERATION_BUFFER)

	master.successfulOperations = 0
	for master.successfulOperations < master.numIntersections {
		set1 := <-master.intersectionChan
		set2 := <-master.intersectionChan
		operation = &Operation{proc, set1, set2}

		worker = <-master.idleWorkerChan
		wg.Add(1)
		go master.runOperation(worker, operation, &wg)
	}

	wg.Wait()

	master.hasMadeFirstTry = true

	if master.successfulOperations == master.numIntersections {
		close(master.failedOperationChan)
		master.hasMadeFirstTry = false
	}

	for operation = range master.failedOperationChan {
		worker = <-master.idleWorkerChan
		wg.Add(1)
		go master.runOperation(worker, operation, &wg)
	}

	wg.Wait()

	// log.Printf("%vx %v operations completed\n", master.successfulOperations, proc)
	return <-master.intersectionChan
}

// runOperation start a single operation on a RemoteWorker and wait for it to return or fail.
func (master *Master) runOperation(remoteWorker *RemoteWorker, operation *Operation, wg *sync.WaitGroup) {

	var (
		err  error
		args *customrpc.RunArgs
	)

	// log.Printf("Running %v (ID: '%v' File: '%v' Worker: '%v')\n", operation.proc, operation.id, operation.filePath, remoteWorker.id)

	args = &customrpc.RunArgs{Set1: operation.set1, Set2: operation.set2}

	// reply := new(struct{})
	reply := customrpc.IntersectReply{}
	err = remoteWorker.callRemoteWorker(operation.proc, args, &reply)
	log.Println("Reply: ", reply)

	if err != nil {
		log.Printf("Operation Failed. Error: %v\n", err)
		wg.Done()
		master.failedWorkerChan <- remoteWorker
		master.failedOperationChan <- operation
	} else {
		wg.Done()
		master.idleWorkerChan <- remoteWorker
		master.successfulOperations++
		if master.hasMadeFirstTry {
			if master.successfulOperations == master.numIntersections {
				close(master.failedOperationChan)
				master.hasMadeFirstTry = false
			}
		}
	}
}
