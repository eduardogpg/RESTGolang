package models

import "errors"

type ValidationError error

var(
  errorUsername = ValidationError(errors.New("El Username no debe de estar vacío."))
  errorShortUsername = ValidationError(errors.New("El Username es demasiado corto."))
  errorLargeUsername = ValidationError(errors.New("El Username es demasiado largo."))
  errorDuplicateUsername = ValidationError(errors.New("El Username ya se encuentra en uso.")) //Para no usar sql

  errorEmail = ValidationError(errors.New("Formato invalido de Email."))

  errorLogin = ValidationError(errors.New("Usuario o password incorrectos."))

  errorPassword = ValidationError(errors.New("El password no debe de estar vacío."))
  errorPasswordEncryption = ValidationError(errors.New("El password no debe de estar vacío."))
)