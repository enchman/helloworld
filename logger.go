package main

import (
	"log"
	"net/http"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)
	l.handler.ServeHTTP(w, r)
}

func AccessLog(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}
