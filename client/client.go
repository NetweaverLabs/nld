package client

import (
	"encoding/gob"
	"log"
	"net"
	"os"

	"github.com/NetwaeversLab/nld/paths"
	"github.com/NetwaeversLab/nld/requests"
)

type Client struct {
	addr    string
	logger  *log.Logger
	encoder *gob.Encoder
	decoder *gob.Decoder
}

// returns a Client struct while will commnuicate with server
func NewClient(host, port string) (*Client, error) {
	// join the host and port to make a addr
	addr := net.JoinHostPort(host, port)
	conn, err := net.Dial("tcp", addr) // dial to the server to get a conn
	if err != nil {
		return nil, err
	}

	// create encoder and decoder
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	// start logger by creating its file in defined path
	lf, err := os.OpenFile(paths.LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	logger := log.New(lf, "NLD CLIENT ", log.Default().Flags())
	return &Client{
		addr:    addr,
		logger:  logger,
		encoder: enc,
		decoder: dec,
	}, nil
}

// Send encodes the request and send its to the server
func (c *Client) Send(req *requests.ServerRequest) error {
	if err := c.encoder.Encode(req); err != nil {
		c.logger.Println("error while encoding request to server: ", err.Error())
		return err
	}
	return nil
}

// Recieve decodes the response from the server, if resp has a nil reference it will panic
func (c *Client) Recieve(resp any) error {
	if resp == nil {
		c.logger.Panicln("resp cannot be nil")
	}
	if err := c.decoder.Decode(resp); err != nil {
		c.logger.Println("error while decoding the response from server: ", err.Error())
		return err
	}
	return nil
}
