package model

type Input struct {
	Microservice string            `json:"ms"`
	Action       string            `json:"action"`
	Params       map[string]string `json:"params"`
	Method       string			   `json:"method"`
}
