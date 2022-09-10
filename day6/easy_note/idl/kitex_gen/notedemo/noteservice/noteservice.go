// Code generated by Kitex v0.3.1. DO NOT EDIT.

package noteservice

import (
	"context"
	"easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return noteServiceServiceInfo
}

var noteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NoteService"
	handlerType := (*notedemo.NoteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateNote": kitex.NewMethodInfo(createNoteHandler, newNoteServiceCreateNoteArgs, newNoteServiceCreateNoteResult, false),
		"MGetNote":   kitex.NewMethodInfo(mGetNoteHandler, newNoteServiceMGetNoteArgs, newNoteServiceMGetNoteResult, false),
		"DeleteNote": kitex.NewMethodInfo(deleteNoteHandler, newNoteServiceDeleteNoteArgs, newNoteServiceDeleteNoteResult, false),
		"QueryNote":  kitex.NewMethodInfo(queryNoteHandler, newNoteServiceQueryNoteArgs, newNoteServiceQueryNoteResult, false),
		"UpdateNote": kitex.NewMethodInfo(updateNoteHandler, newNoteServiceUpdateNoteArgs, newNoteServiceUpdateNoteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "notedemo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notedemo.NoteServiceCreateNoteArgs)
	realResult := result.(*notedemo.NoteServiceCreateNoteResult)
	success, err := handler.(notedemo.NoteService).CreateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceCreateNoteArgs() interface{} {
	return notedemo.NewNoteServiceCreateNoteArgs()
}

func newNoteServiceCreateNoteResult() interface{} {
	return notedemo.NewNoteServiceCreateNoteResult()
}

func mGetNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notedemo.NoteServiceMGetNoteArgs)
	realResult := result.(*notedemo.NoteServiceMGetNoteResult)
	success, err := handler.(notedemo.NoteService).MGetNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceMGetNoteArgs() interface{} {
	return notedemo.NewNoteServiceMGetNoteArgs()
}

func newNoteServiceMGetNoteResult() interface{} {
	return notedemo.NewNoteServiceMGetNoteResult()
}

func deleteNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notedemo.NoteServiceDeleteNoteArgs)
	realResult := result.(*notedemo.NoteServiceDeleteNoteResult)
	success, err := handler.(notedemo.NoteService).DeleteNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceDeleteNoteArgs() interface{} {
	return notedemo.NewNoteServiceDeleteNoteArgs()
}

func newNoteServiceDeleteNoteResult() interface{} {
	return notedemo.NewNoteServiceDeleteNoteResult()
}

func queryNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notedemo.NoteServiceQueryNoteArgs)
	realResult := result.(*notedemo.NoteServiceQueryNoteResult)
	success, err := handler.(notedemo.NoteService).QueryNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceQueryNoteArgs() interface{} {
	return notedemo.NewNoteServiceQueryNoteArgs()
}

func newNoteServiceQueryNoteResult() interface{} {
	return notedemo.NewNoteServiceQueryNoteResult()
}

func updateNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*notedemo.NoteServiceUpdateNoteArgs)
	realResult := result.(*notedemo.NoteServiceUpdateNoteResult)
	success, err := handler.(notedemo.NoteService).UpdateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceUpdateNoteArgs() interface{} {
	return notedemo.NewNoteServiceUpdateNoteArgs()
}

func newNoteServiceUpdateNoteResult() interface{} {
	return notedemo.NewNoteServiceUpdateNoteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest) (r *notedemo.CreateNoteResponse, err error) {
	var _args notedemo.NoteServiceCreateNoteArgs
	_args.Req = req
	var _result notedemo.NoteServiceCreateNoteResult
	if err = p.c.Call(ctx, "CreateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest) (r *notedemo.MGetNoteResponse, err error) {
	var _args notedemo.NoteServiceMGetNoteArgs
	_args.Req = req
	var _result notedemo.NoteServiceMGetNoteResult
	if err = p.c.Call(ctx, "MGetNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest) (r *notedemo.DeleteNoteResponse, err error) {
	var _args notedemo.NoteServiceDeleteNoteArgs
	_args.Req = req
	var _result notedemo.NoteServiceDeleteNoteResult
	if err = p.c.Call(ctx, "DeleteNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest) (r *notedemo.QueryNoteResponse, err error) {
	var _args notedemo.NoteServiceQueryNoteArgs
	_args.Req = req
	var _result notedemo.NoteServiceQueryNoteResult
	if err = p.c.Call(ctx, "QueryNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest) (r *notedemo.UpdateNoteResponse, err error) {
	var _args notedemo.NoteServiceUpdateNoteArgs
	_args.Req = req
	var _result notedemo.NoteServiceUpdateNoteResult
	if err = p.c.Call(ctx, "UpdateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
