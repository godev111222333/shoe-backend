package handler

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/store"
	"github.com/godev111222333/shoe-backend/src/token"
)

type APIServer struct {
	cfg        *misc.APIConfig
	route      *gin.Engine
	store      *store.DbStore
	otpService *OTPService

	tokenMaker token.Maker
}

func NewAPIServer(
	cfg *misc.APIConfig,
	db *store.DbStore,
	otpService *OTPService,
) *APIServer {
	tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		panic(err)
	}

	s := &APIServer{
		cfg:        cfg,
		route:      gin.New(),
		store:      db,
		otpService: otpService,
		tokenMaker: tokenMaker,
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
	authGroup := s.route.Group("/").Use(authMiddleware(s.tokenMaker))

	for _, r := range s.AllRoutes() {
		if !r.RequiredAuth {
			s.route.Handle(r.Method, r.Path, r.Handler)
		} else {
			authGroup.Handle(r.Method, r.Path, r.Handler)
		}
	}
}
