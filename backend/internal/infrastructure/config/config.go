package config
import "os"
type Config struct{HTTPAddr,DatabaseURL,RedisURL,Environment string}
func Load()Config{return Config{HTTPAddr:env("HTTP_ADDR",":8080"),DatabaseURL:env("DATABASE_URL","postgres://vtc:vtc@localhost:5432/vtc?sslmode=disable"),RedisURL:env("REDIS_URL","redis://localhost:6379/0"),Environment:env("APP_ENV","development")}}
func env(k,f string)string{if v:=os.Getenv(k);v!=""{return v};return f}
