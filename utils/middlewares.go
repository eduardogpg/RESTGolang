package utils

import (
  "net/http"
  "fmt"
)

type functionHandler func(w http.ResponseWriter, r *http.Request)

// Un middleware te permiten agregar capas adicionales a la lógica de tu aplicación.

func MiddlewareFunc(function functionHandler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
     function(w, r)
  })
}

func MiddlewareHandle(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
     handler.ServeHTTP(w, r) (w, r)
  })
}
