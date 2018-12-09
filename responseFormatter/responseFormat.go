package responseFormatter

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Datastruct struct {
	//for JSON Formatting
	StatusCode string            `json:"statuscode"`
	Response   interface{}       `json:"response"`
}

func JsonResponse(response interface{}, status string, w http.ResponseWriter) {

	fmt.Println(response)
	data := Datastruct{Response:response, StatusCode:status}
	json.NewEncoder(w).Encode(data)

}