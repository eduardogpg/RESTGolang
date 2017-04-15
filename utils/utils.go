package utils

import "os"

func GetStringEnv(env string, defEnv string) string{
  val := GetEnv(env, defEnv)
  return val.(string)
}

func GetIntEnv(env string, defEnv int) int{
  val := GetEnv(env, defEnv)
  return val.(int)
}

func GetBoolEnv(env string, defEnv bool) bool{
  val := GetEnv(env, defEnv)
  return val.(bool)   
}

func GetEnv(env string, defEnv interface{}) interface{} {
  if val := os.Getenv(env); val == ""{
    return defEnv
  }else{
    return val
  }
}