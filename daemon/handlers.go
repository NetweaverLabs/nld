package daemon

import (
	"github.com/NetwaeversLab/nld/requests"
	"github.com/NetwaeversLab/nld/responses"
)

func Echo(req *requests.DeamonRequest) *responses.DaemonResponse {
	resp := &responses.DaemonResponse{
		Status:  "OK",
		Payload: req.Args,
	}
	return resp
}
