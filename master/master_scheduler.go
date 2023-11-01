package master

import (
	"dii/customrpc"
	"dii/invertedindex"
	"log"
	"sync"
)

// Schedules map operations on remote workers. This will run until InputFilePathChan
// is closed. If there is no worker available, it'll block.
func (master *Master) schedule(task *invertedindex.Task, proc string, filePathChan chan string) int {
	var (
		wg        sync.WaitGroup
		filePath  string
		worker    *RemoteWorker
		operation *Operation
	)

	master.failedOperationChan = make(chan *Operation, RETRY_OPERATION_BUFFER)

	log.Printf("Scheduling %v operations\n", proc)

	// counter = 0
	master.totalNumberOfOperations = 0
	master.successfulOperations = 0
	for filePath = range filePathChan {
		operation = &Operation{proc, master.totalNumberOfOperations, filePath}
		// counter++
		master.totalNumberOfOperations++
		worker = <-master.idleWorkerChan
		wg.Add(1)
		go master.runOperation(worker, operation, &wg)
	}

	wg.Wait()

	master.hasMadeFirstTry = true

	if master.successfulOperations == master.totalNumberOfOperations {
		close(master.failedOperationChan)
		master.hasMadeFirstTry = false
	}

	for operation = range master.failedOperationChan {
		worker = <-master.idleWorkerChan
		wg.Add(1)
		go master.runOperation(worker, operation, &wg)
	}

	wg.Wait()

	log.Printf("%vx %v operations completed\n", master.successfulOperations, proc)
	return master.successfulOperations
}

// runOperation start a single operation on a RemoteWorker and wait for it to return or fail.
func (master *Master) runOperation(remoteWorker *RemoteWorker, operation *Operation, wg *sync.WaitGroup) {

	var (
		err  error
		args *customrpc.RunArgs
	)

	log.Printf("Running %v (ID: '%v' File: '%v' Worker: '%v')\n", operation.proc, operation.id, operation.filePath, remoteWorker.id)

	args = &customrpc.RunArgs{Id: operation.id, FilePath: operation.filePath}
	err = remoteWorker.callRemoteWorker(operation.proc, args, new(struct{}))

	if err != nil {
		log.Printf("Operation %v '%v' Failed. Error: %v\n", operation.proc, operation.id, err)
		wg.Done()
		master.failedWorkerChan <- remoteWorker
		master.failedOperationChan <- operation
	} else {
		wg.Done()
		master.idleWorkerChan <- remoteWorker
		master.successfulOperations++
		if master.hasMadeFirstTry {
			if master.successfulOperations == master.totalNumberOfOperations {
				close(master.failedOperationChan)
				master.hasMadeFirstTry = false
			}
		}
	}
}
