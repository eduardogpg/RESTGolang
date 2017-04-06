package utils

import (
  "net/http"
  "fmt"
  "encoding/json"
)

type Response struct{
  Status      int         `json:"status"`
  Data        interface{} `json:"data"`
  Message     string      `json:"message"`
  contentType string
  write       http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
  return Response{Status: http.StatusOK, write: w, contentType: "application/json" }
}

func SendNotFound(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.NotFound()
  response.Send()
}

func (this *Response) NotFound(){
  this.Status = http.StatusNotFound
  this.Message = "Resource not found."
}

func SendUnprocessableEntity(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.UnprocessableEntity()
  response.Send()
}

func (this *Response) UnprocessableEntity(){
  this.Status = http.StatusUnprocessableEntity
  this.Message = "Unprocessable entity."
}

func SendNoContent(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.UnprocessableEntity()
  response.Send()
}

func (this *Response) NoContent(){
  this.Status = http.StatusNoContent
  this.Message = "No Content."
}

func SendBadRequest(w http.ResponseWriter){
  response := CreateDefaultResponse(w)
  response.BadRequest()
  response.Send()
}

func (this *Response) BadRequest(){
  this.Status = http.StatusBadRequest
  this.Message = "Bad request."
}

func SendData(w http.ResponseWriter, data interface{}){
  response := CreateDefaultResponse(w)
  response.Data = data
  response.Send()
}

func (this *Response) Send() {
  this.write.Header().Set("Content-Type", this.contentType )
  this.write.WriteHeader(this.Status)

  output, _:= json.Marshal(&this)
  fmt.Fprintln(this.write, string(output))
}
