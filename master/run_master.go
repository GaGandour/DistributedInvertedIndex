package master

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
)

// RunMaster will start a master node on the map reduce operations.
// In the distributed model, a Master should serve multiple workers and distribute
// the operations to be executed in order to complete the task.
//   - task: the Task object that contains the mapreduce operation.
//   - hostname: the tcp/ip address on which it will listen for connections.
func RunMaster(hostname string) {
	var (
		query        string
		err          error
		master       *Master
		newRpcServer *rpc.Server
		listener     net.Listener
		// reduceFilePathChan chan string
		// mapOperations      int
	)

	log.Println("Running Master on", hostname)

	master = newMaster(hostname)

	newRpcServer = rpc.NewServer()
	newRpcServer.Register(master)

	if err != nil {
		log.Panicln("Failed to register RPC server. Error:", err)
	}

	master.rpcServer = newRpcServer

	listener, err = net.Listen("tcp", master.address)

	if err != nil {
		log.Panicln("Failed to start TCP server. Error:", err)
	}

	master.listener = listener

	// Start MapReduce Operation

	go master.acceptMultipleConnections()
	go master.handleFailingWorkers()

	for {
		_, err := fmt.Scanln(&query)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		// Schedule workers operations
		//mapOperations = master.schedule(task, "Worker.RunMap", task.InputFilePathChan)

		// Merge the results of multiple II operations into a single result
		// TODO: change function name
		//mergeMapLocal(task, mapOperations)

		// Schedule intersect operations
		// TODO: change function names and implement
		// reduceFilePathChan = fanReduceFilePath(task.NumReduceJobs)
		// reduceOperations = master.schedule(task, "Worker.RunReduce", reduceFilePathChan)

		log.Println("Results found: TODO: print results.")
		log.Println("Time elapsed: TODO: time elapsed.")
	}
	return
}
