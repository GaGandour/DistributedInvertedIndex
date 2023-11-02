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
	Id   int
	set1 []int
	set2 []int
}
