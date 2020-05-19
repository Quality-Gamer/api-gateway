package model

import "github.com/cheekybits/genny/generic"

type JsonResponse struct {
	Response generic.Type `json:"response"`
}

type Response struct {
	Status   string `json:"status"`
	Response map[string]generic.Type `json:"response"`
	Message string `json:"message"`
}
