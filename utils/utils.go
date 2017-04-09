package utils

import(
  "os"
)

func GetEnvOrDefault(env, default string) string{
  if val := os.Getenv(env); val == ""{
    return default
  }else{
    return val
  }
}