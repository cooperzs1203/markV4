/**
* @Author: Cooper
* @Date: 2019/12/13 20:08
 */

package mnet

import (
	"log"
	"markV4/mface"
)

func NewServer() (mface.MServer, error) {
	config := defaultConfig()
	return NewServerWithConfig(config)
}

func NewServerWithConfig(config mface.MConfig) (mface.MServer, error) {
	s := &server{
		status:         Serve_Status_UnStarted,
		config:         config,
		connManager:    newConnManager(),
		messageManager: newMessageManager(),
		routeManager:   newRouteManager(),
	}

	s.connManager.SetServer(s)
	s.messageManager.SetServer(s)
	s.routeManager.SetServer(s)

	err := s.Load()

	return s, err
}

type server struct {
	status int

	config         mface.MConfig
	connManager    mface.MConnManager
	messageManager mface.MMessageManager
	routeManager   mface.MRouteManager
}

func (s *server) Load() error {
	if err := s.config.Load(); err != nil {
		return err
	}

	if err := s.connManager.Load(); err != nil {
		return err
	}

	if err := s.messageManager.Load(); err != nil {
		return err
	}

	if err := s.routeManager.Load(); err != nil {
		return err
	}

	// todo:server load
	log.Printf("[Server] Load")
	s.status = Serve_Status_Load

	return nil
}

func (s *server) Start() error {
	if err := s.connManager.Start(); err != nil {
		return err
	}

	if err := s.messageManager.Start(); err != nil {
		return err
	}

	if err := s.routeManager.Start(); err != nil {
		return err
	}

	log.Printf("[Server] Start")
	s.status = Serve_Status_Start

	return nil
}

func (s *server) Reload() error {
	if err := s.config.Reload(); err != nil {
		return err
	}

	if err := s.connManager.Reload(); err != nil {
		return err
	}

	if err := s.messageManager.Reload(); err != nil {
		return err
	}

	if err := s.routeManager.Reload(); err != nil {
		return err
	}

	// todo:server reload
	log.Printf("[Server] Reload")
	s.status = Serve_Status_Reload

	return nil
}

func (s *server) Status() int {
	return s.status
}

func (s *server) Stop() error {
	log.Printf("[Server] Start Ending")
	s.status = Serve_Status_Ending

	if err := s.connManager.StartEnding(); err != nil {
		return err
	}

	if err := s.messageManager.StartEnding(); err != nil {
		return err
	}

	if err := s.routeManager.StartEnding(); err != nil {
		return err
	}

	if err := s.routeManager.OfficialEnding(); err != nil {
		return err
	}

	if err := s.messageManager.OfficialEnding(); err != nil {
		return err
	}

	if err := s.connManager.OfficialEnding(); err != nil {
		return err
	}

	s.status = Serve_Status_Stopped
	log.Printf("[Server] Stopped")
	return nil
}
