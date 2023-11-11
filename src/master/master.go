package master

import (
	"dii/invertedindex"
	"log"
	"net"
	"net/rpc"
	"sync"
)

const (
	IDLE_WORKER_BUFFER     = 100
	RETRY_OPERATION_BUFFER = 100
	DOCS_PATH              = "../books/pg"
)

type Master struct {
	// Network
	address   string
	rpcServer *rpc.Server
	listener  net.Listener

	// Workers handling
	workersMutex sync.Mutex
	workers      map[int]*RemoteWorker
	totalWorkers int // Used to generate unique ids for new workers

	idleWorkerChan chan *RemoteWorker

	// Inverted index
	ii               invertedindex.InvertedIndex
	intersectionChan chan []int

	// Fault Tolerance
	numIntersections     int
	totalOperations      int
	successfulOperations int
}

type Operation struct {
	// id   int
	proc string
	set1 []int
	set2 []int
}

// Construct a new Master struct
func newMaster(address string) (master *Master) {
	master = new(Master)
	master.address = address
	master.workers = make(map[int]*RemoteWorker, 0)
	master.idleWorkerChan = make(chan *RemoteWorker, IDLE_WORKER_BUFFER)
	// TODO: BUGA
	master.ii = invertedindex.BuildInvertedIndex(DOCS_PATH)
	// master.ii = invertedindex.InvertedIndex{}
	master.totalWorkers = 0
	master.successfulOperations = 0
	master.totalOperations = 0
	master.numIntersections = 0
	return
}

// acceptMultipleConnections will handle the connections from multiple workers.
func (master *Master) acceptMultipleConnections() {
	var (
		err     error
		newConn net.Conn
	)

	log.Printf("Accepting connections on %v\n", master.listener.Addr())

	for {
		newConn, err = master.listener.Accept()

		if err == nil {
			go master.handleConnection(&newConn)
		} else {
			log.Println("Failed to accept connection. Error: ", err)
			break
		}
	}

	log.Println("Stopped accepting connections.")
}

// Handle a single connection until it's done, then closes it.
func (master *Master) handleConnection(conn *net.Conn) error {
	master.rpcServer.ServeConn(*conn)
	(*conn).Close()
	return nil
}
