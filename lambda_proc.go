package lambda_proc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

type (
	Handler func(*Context, json.RawMessage) (interface{}, error)

	Context struct {
		AwsRequestID             string `json:"awsRequestId"`
		FunctionName             string `json:"functionName"`
		FunctionVersion          string `json:"functionVersion"`
		Invokeid                 string `json:"invokeid"`
		IsDefaultFunctionVersion bool   `json:"isDefaultFunctionVersion"`
		LogGroupName             string `json:"logGroupName"`
		LogStreamName            string `json:"logStreamName"`
		MemoryLimitInMB          string `json:"memoryLimitInMB"`
		EventCounter             int    `json:"-"`
	}

	Payload struct {
		// custom event fields
		Event json.RawMessage `json:"event"`

		// default context object
		Context *Context `json:"context"`
	}

	ErrorResponse struct {
		// Any errors that occur during processing
		// or are returned by handlers are returned
		Error string `json:"error"`
	}
)

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error: err.Error(),
	}
}

var eventCounter int // process event counter

func Run(handler Handler) {
	RunStream(handler, os.Stdin, os.Stdout)
}

func RunStream(handler Handler, Stdin io.Reader, Stdout io.Writer) {

	stdin := json.NewDecoder(Stdin)
	stdout := json.NewEncoder(Stdout)

	for ; ; eventCounter++ {
		if err := func() (err error) {
			defer func() {
				if e := recover(); e != nil {
					err = fmt.Errorf("panic: %v", e)
				}
			}()
			var payload Payload
			if err := stdin.Decode(&payload); err != nil {
				return err
			}
			payload.Context.EventCounter = eventCounter
			data, err := handler(payload.Context, payload.Event)
			if err != nil {
				return err
			}
			return stdout.Encode(data)
		}(); err != nil {
			if encErr := stdout.Encode(NewErrorResponse(err)); encErr != nil {
				// bad times
				logrus.Fatalf("Failed to encode err response!", encErr.Error())
			}
		}
	}
}
