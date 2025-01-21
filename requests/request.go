package requests

type ServerRequest struct {
	Type    string `gob:"Type"`
	Payload string `gob:"Payload"`
	Time    string `gob:"Time"`
}

type DeamonRequest struct {
	Cmd  string `gob:"Cmd"`
	Args any    `gob:"Args"`
}
