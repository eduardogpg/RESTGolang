package models

import "errors"

type ValidationError error

var(
  errorUsername = ValidationError(errors.New("El Username no debe de estar vacío"))
  errorShortUsername = ValidationError(errors.New("El Username es demasiado corto"))
  errorLargeUsername = ValidationError(errors.New("El Username es demasiado largo"))

  errorEmail = ValidationError(errors.New("Formato invalido de Email"))

  errorEmail = ValidationError(errors.New("El password no debe de estar vacío"))
)