// Copyright 2015 The go-symverse Authors
// This file is part of the go-symverse library.
//
// The go-symverse library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-symverse library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-symverse library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewRpcRequest(id json.RawMessage, method string, param json.RawMessage) *jsonRequest {
	jsonRpcRequest := &jsonRequest{
		Version: jsonrpcVersion,
		Id:      id,
		Method:  method,
		Payload: param,
	}
	return jsonRpcRequest
}

func RequestJSON(c *gin.Context) {
	var in *jsonRequest
	var response interface{}

	//request body를 json으로 바인딩
	if err := c.ShouldBindJSON(&in); err != nil {
		response = CreateErrorResponse(in.Id, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//just reflect the input
	response = CreateResponse(in.Id, &in)

	c.JSON(http.StatusOK, response)

}

func convertParam(method string, Payload json.RawMessage) error {
	//temporary string
	//it plan to chage depend on the method come in.
	var str string
	if err := json.Unmarshal(Payload, &str); err == nil {
		return nil
	}
	return fmt.Errorf("invalid request param")

}

// checkReqId returns an error when the given reqId isn't valid for RPC method calls.
// valid id's are strings, numbers or null
func checkReqId(reqId json.RawMessage) error {
	if len(reqId) == 0 {
		return fmt.Errorf("missing request id")
	}
	if _, err := strconv.ParseFloat(string(reqId), 64); err == nil {
		return nil
	}
	var str string
	if err := json.Unmarshal(reqId, &str); err == nil {
		return nil
	}
	return fmt.Errorf("invalid request id")
}
