// common_rpc.go defined all the parameters used in RPC between
// master and workers
package customrpc

type RegisterArgs struct {
	WorkerHostname string
}

type RegisterReply struct {
	WorkerId int
}

type RunArgs struct {
	// Id   int
	Set1 []int
	Set2 []int
}

type IntersectReply struct {
	Result []int
}
