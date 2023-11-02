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

		// Split into words

		// Create chanel of int lists -> retrieval results

		// (FAZER ISSO NO SCHEDULE:)
		// for loop infinito
		// 	- tentar pegar n elementos do canal (n=2)
		// 	- mandar n elementos pra intersect (worker)
		//  - pega o resultado e poe no canal

		// o resultado final é uma lista de indices

		// pegar o endereco desses indices
		// for result in resultado:
		// ii.docs[result]

		// printar os enderecos

		// Schedule intersect operations
		// TODO: change function names and implement
		// reduceFilePathChan = fanReduceFilePath(task.NumReduceJobs)
		// reduceOperations = master.schedule(task, "Worker.RunIntersect", reduceFilePathChan)

		log.Println("Results found: TODO: print results.")
		log.Println("Time elapsed: TODO: time elapsed.")
	}
	return
}

// implementar a mensagem de volta do worker / escrever txt (dificil)
// implementar o schedule (dificil)

// implementar a função de intersect (ez)
// fazer o master ler o inverted index na inicialização (ez)
// indexar esses txts (medio)

// fazer o dockerfile (ez) (segundo o copilot)
