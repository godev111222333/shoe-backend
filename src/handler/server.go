package handler

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/store"
)

type APIServer struct {
	cfg        *misc.APIConfig
	route      *gin.Engine
	store      *store.DbStore
	otpService *OTPService
}

func NewAPIServer(
	cfg *misc.APIConfig,
	db *store.DbStore,
	otpService *OTPService,
) *APIServer {
	s := &APIServer{
		cfg:        cfg,
		route:      gin.New(),
		store:      db,
		otpService: otpService,
	}

	s.setUp()
	return s
}

func (s *APIServer) Run() error {
	fmt.Printf("API server running at port: %s\n", s.cfg.Port)

	return s.route.Run(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
}

func (s *APIServer) setUp() {
	s.registerMiddleware()
	s.registerHandlers()
}

func (s *APIServer) registerMiddleware() {
	s.route.Use(cors.Default())
}

func (s *APIServer) registerHandlers() {
	for _, r := range s.AllRoutes() {
		s.route.Handle(r.Method, r.Path, r.Handler)
	}
}
