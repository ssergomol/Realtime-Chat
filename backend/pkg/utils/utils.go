package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Couldn't read request body")
		return
	}

	err = json.Unmarshal([]byte(body), x)

	if err != nil {
		fmt.Printf("Couldn't unmarshal json data from request")
		return
	}
}

func SetBodyInfoMessage(writer http.ResponseWriter, message string) {
	data := make(map[string]string, 1)
	data["message"] = string(message)
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Couldn't marshal information message\n")
	}
	writer.Write(res)
}
