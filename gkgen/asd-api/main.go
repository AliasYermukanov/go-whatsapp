package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	
	"github.com/AliasYermukanov/gkgen/asd-api/src/config"
	"github.com/AliasYermukanov/gkgen/asd-api/src/asd"
	
)


var s asd.Service

func main() {

	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address only port :8080")
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	logr := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.Level(5),
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
	asd.Loger = logr

	err := config.GetConfigs()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	s = asd.NewService()
	s = asd.NewLoggingService(log.With(logger, "component", "asd"), s)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()
	mux.Handle("/v1/asd/", asd.MakeHandler(s, httpLogger))
	http.Handle("/v1/", accessControl(mux))
	http.HandleFunc("/check", config.Healthchecks)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening showcase-api")
		errs <- http.ListenAndServe(*httpAddr, mux)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

type serializedLogger struct {
	mtx sync.Mutex
	log.Logger
}

func (l *serializedLogger) Log(keyvals ...interface{}) error {
    l.mtx.Lock()
    defer l.mtx.Unlock()
    return l.Logger.Log(keyvals...)
}
