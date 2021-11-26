package homie

import (
	"fmt"
	"net/http"

	"github.com/SkurkovPavel/Homie/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Service struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func NewService(config *Config, logger *logrus.Logger, router *mux.Router) *Service {
	return &Service{
		config: config,
		logger: logger,
		router: router,
	}
}

func (s *Service) Start() error {

	if err := s.configureLogger(); err != nil {
		return fmt.Errorf("could not congure logger: %s", err.Error())
	}

	s.configureRouter()

	if err := s.configureStorage(); err != nil {
		return err
	}

	s.logger.Infof("server start saccessful")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Service) configureRouter() {
	s.router.HandleFunc("/alice", s.AliceCommander(false))
}
func (s *Service) configureStorage() error {

	st := storage.NewStorage(s.config.Storage)

	if err := st.Open(); err != nil {
		return err
	}

	s.storage = st
	return nil
}
