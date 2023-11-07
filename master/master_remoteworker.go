package master

import (
	"dii/customrpc"
	"fmt"
	"net/rpc"
)

type workerStatus string

const (
	WORKER_IDLE    workerStatus = "idle"
	WORKER_RUNNING workerStatus = "running"
)

type RemoteWorker struct {
	id       int
	hostname string
	status   workerStatus
}

// Call a RemoteWork with the procedure specified in parameters. It will also handle connecting
// to the server and closing it afterwards.
func (worker *RemoteWorker) callRemoteWorker(proc string, args interface{}, reply *customrpc.IntersectReply) error {
	var (
		err    error
		client *rpc.Client
	)

	client, err = rpc.Dial("tcp", worker.hostname)

	if err != nil {
		return err
	}

	defer client.Close()
	err = client.Call(proc, args, reply)

	if err != nil {
		var tmpClient *rpc.Client
		tmpClient, err = rpc.Dial("tcp", worker.hostname)
		defer tmpClient.Close()
		if err != nil {
			return err
		}

		err = tmpClient.Call(proc, args, reply)

		return err
	}

	return nil
}
