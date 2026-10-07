package main

import("log";"net/http";"time";"github.com/armelo10/vtc_go/backend/internal/infrastructure/config";"github.com/armelo10/vtc_go/backend/internal/interfaces/httpapi")
func main(){cfg:=config.Load();srv:=&http.Server{Addr:cfg.HTTPAddr,Handler:httpapi.NewServer(cfg).Handler(),ReadHeaderTimeout:5*time.Second};log.Printf("vtc api listening on %s",cfg.HTTPAddr);if err:=srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{log.Fatal(err)}}
