package request

import (
	"encoding/json"
	"fmt"
	"gateway/conf"
	"gateway/database"
	"gateway/model"
	"github.com/cheekybits/genny/generic"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"net/http"
)

func Route(c echo.Context) (err error) {
	input := new(model.Input)

	if err = c.Bind(input); err != nil {
		return
	}

	var res model.JsonResponse
	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)

	if len(input.Microservice) == 0 {
		var r model.Response
		r.Status = conf.ErrorCode
		r.Message = conf.ErrorInputMessage
		res.Response = r
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(res)
	}

	microservice := conf.Microservice
	action := ""
	ms := input.Microservice

	if len(input.Action) > 0 {
		action = input.Action
	}

	incrRequest(input.Microservice,input.Action)

	var response generic.Type

	if input.Method == "POST" {
		response = makePOSTRequest(microservice[ms], input.Params, action, ms)
	} else {
		response = makeGETRequest(microservice[ms], input.Params, action, ms)
	}

	res.Response = response
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(res)
}

func makePOSTRequest(url string, params map[string]string, action string,ms string) generic.Type  {
	urlComplete := url + "/" + action
	var strPost = []byte("POST")
	var strRequestURI = []byte(urlComplete)
	paramsJSON, _ := json.Marshal(params)

	req := fasthttp.AcquireRequest()
	req.SetBody(paramsJSON)
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	req.Header.Set("Content-Type","application/json")
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		stopMicroservice(ms)
		var e model.Response
		e.Status = conf.ErrorCode
		e.Message = conf.MSOffMessage
		return e
	}
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	turnOnMicroservice(ms)

	body := res.Body()
	var r generic.Type
	json.Unmarshal(body,&r)

	return r
}

func makeGETRequest(url string, params map[string]string, action string,ms string) generic.Type  {
	urlComplete := url + "/" + action
	var strPost = []byte("GET")

	if len(params) > 0 {
		urlComplete += "?"

		for key,value := range params {
			urlComplete += key + "="
			urlComplete += value + "&"
		}

		size := len(urlComplete)
		urlComplete = urlComplete[:size-1]
	}

	var strRequestURI = []byte(urlComplete)

	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		stopMicroservice(ms)
		var e model.Response
		e.Status = conf.ErrorCode
		e.Message = conf.MSOffMessage
		return e
	}
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	turnOnMicroservice(ms)

	body := res.Body()
	var r generic.Type
	json.Unmarshal(body,&r)

	return r
}

func incrRequest(ms,action string) {
	keyMS := conf.GetMicroserviceKeyCount(ms)
	keyAction := conf.GetMicroserviceActionKeyCount(ms,action)

	database.IncrValue(keyMS)
	database.IncrValue(keyAction)
}

func stopMicroservice(ms string) {
	database.HSetKey(conf.GetStoppedMicroserviceKey(),ms,ms)
}

func turnOnMicroservice(ms string) {
	database.HDelField(conf.GetStoppedMicroserviceKey(),ms)
}

func debugStoppedMicroservices(){
	ms := database.HValKey(conf.GetStoppedMicroserviceKey())
	fmt.Print(ms)
}