package server

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/Crypto/cryptoapp-data-collection/misc"
)

type ServerInfo struct {
	Server    string `json:"server"`
	Hostname  string `json:"hostname"`
	Environment string `json:"environment"`
}

type Server struct {
	Info      *ServerInfo
	Uptime    time.Time
	Server    *http.Server
	Producer  *Loader
}

func (s *Server) Run(conf *misc.Http, address string) error {
	router := s.NewRouter()
	s.Server = &http.Server{
		Addr:              address,
		ReadTimeout:       time.Duration(conf.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(conf.WriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(conf.ReadTimeout*2) * time.Second,
		Handler:           router,
	}

	log.Info("HTTP server ready", "address", conf.Address[1:])

	return s.Server.ListenAndServe()
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if s.Server != nil {
		if err := s.Server.Shutdown(ctx); err != nil {
			log.Error("Failed to shut down http server cleanly", "error", err)

			// Close all open connections
			if err := s.Server.Close(); err != nil {
				log.Error("Failed to force-close http server", "error", err)
			}
		}
	}

	return nil
}

func (s *Server) NewRouter() *mux.Router {
	return AddRouter(s)
}

func AddRouter(s *Server) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//s.addRoute(router, "Index", "GET", "/", s.Index)

	return router
}

func (s *Server) addRoute(router *mux.Router, name string, method string, pattern string, fn http.HandlerFunc) {
	var handler http.Handler
	handler = s.WrapRequest(fn, name)

	router.Methods(method).Path(pattern).Name(name).Handler(handler)

}

type LoggedWriter struct {
	http.ResponseWriter

	statusCode int
}

func NewLoggedWriter(w http.ResponseWriter) *LoggedWriter {
	return &LoggedWriter{w, http.StatusOK}
}


func (s *Server) WrapRequest(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := NewLoggedWriter(w)
		inner.ServeHTTP(lw, r)

		duration := time.Since(start)
		statusCode := lw.statusCode

		log.Info(name,"method", r.Method,"uri", r.RequestURI,"name",name,"status",statusCode,"duration",duration)
	})
}