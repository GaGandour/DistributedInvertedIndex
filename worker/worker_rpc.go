package worker

import (
	"dii/customrpc"
	"log"
)

// RPC - RunIntersect
// Run the retrieval operation defined in the task and return when it's done.
func (worker *Worker) RunIntersect(args *customrpc.RunArgs, _ *struct{}) ([]int, error) {
	// TODO: Implement this method.
	// Should retrieve the data given the query.
	return []int{}, nil
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
