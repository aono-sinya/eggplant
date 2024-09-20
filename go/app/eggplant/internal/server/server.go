package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olahol/melody"
	"log"
	"os"
)

type Server struct {
	echo   *echo.Echo
	melody *melody.Melody
	Logger *log.Logger
}

var instance = &Server{}

func GetInstance() *Server {
	return instance
}

func Initialize() error {
	return GetInstance().initialize()
}

func Start() error {
	return GetInstance().start()
}

func HandleGet(url string, f func(*Server, echo.Context) error) {
	GetInstance().handleGet(url, f)
}

func HandlePost(url string, f func(*Server, echo.Context) error) {
	GetInstance().handlePost(url, f)
}

func (s *Server) initialize() error {
	file, err := os.OpenFile("/app/log/crane.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	s.echo = echo.New()
	s.melody = melody.New()
	s.Logger = log.New(file, "", log.LstdFlags)
	return nil
}

func (s *Server) start() error {
	s.Logger.Print("Starting server")
	if s.echo == nil {
		return fmt.Errorf("web server not create")
	}
	if s.melody == nil {
		return fmt.Errorf("websocket server not available")
	}
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	s.melody.HandleConnect(func(ms *melody.Session) {})
	s.melody.HandleMessage(func(ms *melody.Session, msg []byte) {})
	s.melody.HandleDisconnect(func(ms *melody.Session) {})

	errChan := make(chan error, 1)
	go func(addr string) {
		errChan <- s.echo.Start(addr)
	}(":7777")
	err := <-errChan
	return err
}

func (s *Server) handleGet(url string, f func(*Server, echo.Context) error) {
	s.echo.GET(url, func(ctx echo.Context) error {
		return f(s, ctx)
	})
}

func (s *Server) handlePost(url string, f func(*Server, echo.Context) error) {
	s.echo.POST(url, func(ctx echo.Context) error {
		return f(s, ctx)
	})
}
