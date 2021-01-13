package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"roberto.local/echo/model"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	port := os.Getenv("PORT")

	log.Infoln("")
	log.Infoln("Waiting for requests on port "+port+"...")

	// set the endpoints
	r.GET("/info", Info)
	r.GET("/health", Health)
	r.GET("/metrics", Metrics(promhttp.Handler()))
	r.POST("/echo", Echo)

	// start the server
	HandleFatalErrors(http.ListenAndServe(":"+port, r))
}

func Metrics(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func Echo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	funcName := "Echo()"

	log.Infoln("")
	log.Infoln("echoing back...")
	log.Infoln("")

	// get the body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	response := CheckError(funcName,err)
	body := string(bodyBytes)

	// create the request object using the json input
	var request model.EchoRequest
	if response.Result == SuccessResult {
		response = CheckError(funcName,json.Unmarshal([]byte(body), &request))
	}
	if response.Result == SuccessResult {
		response.Value = request.Request // echoing back...
	}

	// convert response object into json to be sent back
	responseAsJSON, _ := json.Marshal(response)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// log output of the response
	log.Debugln("")
	if response.Result == SuccessResult {
		log.Infoln("response: " + string(responseAsJSON))
	}else{
		log.Errorln("response: " + string(responseAsJSON))
	}

	_,err = fmt.Fprintf(w, "%s", responseAsJSON); HandleFatalErrors(err)
}

const (
	AppName = "echo-go"
	AppDescription = "Simple application used to test connectivity between APIs"
	SuccessResult = "SUCCESS"
	FailureResult = "FAILURE"
)

func HandleFatalErrors(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckError(callingFunction string,err error) model.StandardResponse {
	response := model.StandardResponse{
		Description: AppName,
		Result:      SuccessResult,
		Value:       callingFunction+" succeeded",
	}

	if err != nil {
		response.Result = FailureResult
		response.Value = callingFunction+" failed with error: "+err.Error()
	}

	return response
}

func Health(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infoln("")
	log.Infoln("getting the app info...")
	log.Infoln("")

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	log.Debugln("result: OK")
	fmt.Fprintf(w, "%s", "OK") // TODO: do a more thorough check to come with this answer
}

func Info(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infoln("")
	log.Infoln("getting the app info...")
	log.Infoln("")

	// struct/object to store the info
	i := model.Info{
		Name:        AppName,
		Description: AppDescription,
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(i)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	log.Debugln("result: " + string(uj))
	fmt.Fprintf(w, "%s", uj) // Fprintf is sending uj (formatted response) to w's (http response writer) pointer
}

