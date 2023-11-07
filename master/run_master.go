package master

import (
	"bufio"
	"log"
	"net"
	"net/rpc"
	"os"
	"strings"
	"time"
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

	go master.acceptMultipleConnections()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		query = scanner.Text()
		start := time.Now()

		// Split into words
		words := strings.Split(query, " ")
		master.numIntersections = len(words) - 1

		// Make retrieval to all words in word
		master.intersectionChan = make(chan []int, master.numIntersections+1)
		for _, word := range words {
			master.intersectionChan <- master.ii.Retrieve(word)
		}

		results := master.schedule("Worker.RunIntersect")

		// o resultado final é uma lista de indices
		end := time.Now()
		log.Println("Results found:")
		for _, result := range results {
			// log.Println(result)
			log.Println(master.ii.Docs[result])
		}

		log.Printf("Time elapsed: %s\n", end.Sub(start))
		close(master.intersectionChan)
	}
}

// fazer o master ler o inverted index na inicialização (ez)
// indexar esses txts (medio)

// fazer o dockerfile (ez) (segundo o copilot)
