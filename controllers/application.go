package controllers

import (
  "../utils"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
  context := make(map[string]interface{})
  context["Authenticated"] = utils.IsAuthenticated(r)
  
  utils.RenderTemplate(w, "application/index", context)
}
