package responses

type DaemonResponse struct {
	Status  string   `gob:"Status"`
	Payload []string `gob:"Payload"`
}
