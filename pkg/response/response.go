package response

import (
	"encoding/json"
	"log"
	"net/http"
)

//return not found 404
func NotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

//return bad request 400
func BadRequest(w http.ResponseWriter, err *ErrResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	resp, _ := json.Marshal(err)
	w.Write(resp)
}

//return bad request 400 if invalid decoded request body
func RenderInvalidDecodeRequestBody(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	resp, _ := json.Marshal(ErrInvalidRequest(InvalidDecodeRequestBody))
	w.Write(resp)
}

//return get 200
func Ok(w http.ResponseWriter, payload interface{}) {
	data := &baseResponse{
		Data: payload,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

//return post - 201
func Created(w http.ResponseWriter, payload interface{}) {
	data := &baseResponse{
		Data: payload,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

//return delete - 204
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

//output response success
type baseResponse struct {
	Data interface{} `json:"data"`
}
