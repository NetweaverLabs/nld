package daemon

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/NetweaverLabs/nld/paths"
	"github.com/NetweaverLabs/nld/requests"
	"github.com/NetweaverLabs/nld/responses"
)

type Daemon struct {
	logger  *log.Logger
	handler map[string]DaemonHandler
}

// returns a daemon struct which will communicate with the cli
func NewDaemon() (*Daemon, error) {
	lf, err := os.OpenFile(paths.LOGFILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	logger := log.New(lf, "DAEMON ", log.Default().Flags())
	return &Daemon{
		logger:  logger,
		handler: make(map[string]DaemonHandler),
	}, nil
}

func (d *Daemon) AllHandlers() {
	d.handler["echo"] = Echo
	d.handler["create"] = Create
}

func (d *Daemon) Start() {
	if _, err := os.Stat(paths.UNIXSOCKET); err == nil {
		d.logger.Println("cleaning the socket")
		os.Remove(paths.UNIXSOCKET)
	}
	d.AllHandlers()
	d.logger.Println("loaded all handlers")
	ln, err := net.Listen("unix", paths.UNIXSOCKET)
	if err != nil {
		d.logger.Println("error starting listner: ", err.Error())
	}
	defer ln.Close()
	d.logger.Println("listner is ready for cli to connect")
	for {
		conn, err := ln.Accept()
		if err != nil {
			d.logger.Println("error accepting connection: ", err.Error())
			return
		}
		defer conn.Close()
		d.logger.Println("connected successfully with cli")
		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)
		d.logger.Println("encoder and decoder are ready")
		for {
			req := &requests.DeamonRequest{}
			if err := decoder.Decode(req); err != nil {
				if err == io.EOF {
					d.logger.Println("client has closed the connection")
					break
				} else {
					d.logger.Println("error while decoding request: ", err.Error())
				}
			}
			fmt.Println("req:", req)
			resp := &responses.DaemonResponse{}
			if h, ok := d.handler[req.Cmd]; ok {
				resp = h(req)
			} else {
				resp.Status = "NOTOK"
				resp.Payload = "command not found in daemon"
			}
			if err := encoder.Encode(resp); err != nil {
				d.logger.Println("error while encoding the response to cli: ", err.Error())
			}
			d.logger.Println("a response is sent with status: ", resp.Status)

		}
	}
}
