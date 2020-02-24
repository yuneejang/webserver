package rpc

import (
	"encoding/json"
	"fmt"
)

const (
	jsonrpcVersion           = "2.0"
	serviceMethodSeparator   = "_"
	subscribeMethodSuffix    = "_subscribe"
	unsubscribeMethodSuffix  = "_unsubscribe"
	notificationMethodSuffix = "_subscription"
)

type jsonRequest struct {
	Method  string          `json:"method"`
	Version string          `json:"jsonrpc"`
	Id      json.RawMessage `json:"id,omitempty"`
	Payload json.RawMessage `json:"params,omitempty"`
}

type jsonSuccessResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonErrResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Error   jsonError   `json:"error"`
}

type Error interface {
	Error() string  // returns the message
	ErrorCode() int // returns the code
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}

func (err *jsonError) ErrorCode() int {
	return err.Code
}

// CreateResponse will create a JSON-RPC success response with the given id and reply as result.
func CreateResponse(id interface{}, reply interface{}) interface{} {
	return &jsonSuccessResponse{Version: jsonrpcVersion, Id: id, Result: reply}
}

// CreateErrorResponse will create a JSON-RPC error response with the given id and error.
func CreateErrorResponse(id interface{}, err string) interface{} {
	return &jsonErrResponse{Version: jsonrpcVersion, Id: id, Error: jsonError{Message: err}}
}

// // CreateResponse will create a JSON-RPC success response with the given id and reply as result.
// func CreateResponse(id interface{}, reply interface{}) interface{} {
// 	return &jsonSuccessResponse{Version: jsonrpcVersion, Id: id, Result: reply}
// }

// // CreateErrorResponse will create a JSON-RPC error response with the given id and error.
// func CreateErrorResponse(id interface{}, err Error) interface{} {
// 	return &jsonErrResponse{Version: jsonrpcVersion, Id: id, Error: jsonError{Code: err.ErrorCode(), Message: err.Error()}}
// }
