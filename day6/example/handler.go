package main

import (
	"context"
	"example/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return &api.Response{Message: req.Message}, nil
}

// Add implements the EchoImpl interface.
func (s *EchoImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
// TODO: Your code here...
	
return &api.Response{Sum:req.First+req.Second}
}
