package types

import (
	"github.com/NetwaeversLab/nld/requests"
	"github.com/NetwaeversLab/nld/responses"
)

type DaemonHandler func(*requests.DeamonRequest) *responses.DaemonResponse
