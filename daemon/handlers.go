package daemon

import (
	"github.com/NetweaverLabs/nld/requests"
	"github.com/NetweaverLabs/nld/responses"
)

type DaemonHandler func(req *requests.DeamonRequest) *responses.DaemonResponse

func Echo(req *requests.DeamonRequest) *responses.DaemonResponse {
	resp := &responses.DaemonResponse{
		Status:  "OK",
		Payload: req.Args,
	}
	return resp
}

func Create(req *requests.DeamonRequest) *responses.DaemonResponse {
	resp := &responses.DaemonResponse{
		Status:  "OK",
		Payload: req.Args,
	}
	return resp
}
