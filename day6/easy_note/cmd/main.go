package main

import (
	"net"
	"src/ByteDance/day6/easy_note/kitex_gen/notedemo/noteservice"
	"src/ByteDance/day6/easy_note/pkg/constants"
	"src/ByteDance/day6/easy_note/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.NoteServiceName)
	rpc.InitRPC()
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()
	svr := noteservice.NewServer(new(noteserviceimpl))
}
