package worker

import (
	"dii/customrpc"
	"dii/invertedindex"
	"log"
)

// RPC - RunIntersect
func (worker *Worker) RunIntersect(args *customrpc.RunArgs, reply *customrpc.IntersectReply) error {
	set1 := args.Set1
	set2 := args.Set2
	reply.Result = invertedindex.Intersect(set1, set2)
	return nil
}

// RPC - Done
// Will be called by Master when the task is done.
func (worker *Worker) Done(_ *struct{}, _ *struct{}) error {
	log.Println("Done.")
	defer func() {
		close(worker.done)
	}()
	return nil
}
