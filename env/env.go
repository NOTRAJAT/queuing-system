package env

import (
	"os"
	"strconv"
)

type Config struct {
	Redis_Password string
	Redis_Host_PORT string
	Redis_DB int
	

}

var Env = initconfig()

func initconfig() Config {
	return Config{
		Redis_Password: getOSENV("REDIS_PASSWORD",""),
		Redis_Host_PORT: getOSENV("REDIS_HOST","0.tcp.in.ngrok.io:13057"),
		Redis_DB: func()int{
			value,_:= strconv.Atoi( getOSENV("REDIS_DB","0"))
			return value
		}(),
	
	}
}

func getOSENV(key string, fallback string)string {
	if value,ok:=os.LookupEnv(key);ok{
		return value
	}
	return fallback

}