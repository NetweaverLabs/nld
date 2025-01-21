package responses

type DaemonResponse struct {
	Status  string `gob:"Status"`
	Payload any    `gob:"Payload"`
}
