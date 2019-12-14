/**
* @Author: Cooper
* @Date: 2019/12/13 20:24
 */

package mnet

import (
	"log"
	"markV4/mface"
)

func newMessageManager() mface.MMessageManager {
	mm := &messageManager{
		status:       Serve_Status_UnStarted,
		server:       nil,
		requestChan:  newChannel(),
		responseChan: newChannel(),
	}
	return mm
}

type messageManager struct {
	status int
	server mface.MServer

	requestChan  mface.MChannel
	responseChan mface.MChannel
}

func (mm *messageManager) SetServer(s mface.MServer) {
	mm.server = s
}

func (mm *messageManager) Load() error {
	log.Printf("[MessageManager] Load")
	mm.status = Serve_Status_Load

	mm.requestChan.SetSize(mm.server.Config().MMReqCS())
	if err := mm.requestChan.Load(); err != nil {
		return err
	}

	mm.responseChan.SetSize(mm.server.Config().MMRspCS())
	if err := mm.responseChan.Load(); err != nil {
		return err
	}

	return nil
}

func (mm *messageManager) Start() error {
	log.Printf("[MessageManager] Start")
	mm.status = Serve_Status_Start

	if err := mm.requestChan.Start(); err != nil {
		return err
	}

	if err := mm.responseChan.Start(); err != nil {
		return err
	}

	return nil
}

func (mm *messageManager) Reload() error {
	log.Printf("[MessageManager] Reload")
	mm.status = Serve_Status_Reload

	mm.requestChan.SetSize(mm.server.Config().MMReqCS())
	if err := mm.requestChan.Reload(); err != nil {
		return err
	}

	mm.responseChan.SetSize(mm.server.Config().MMRspCS())
	if err := mm.responseChan.Reload(); err != nil {
		return err
	}

	return nil
}

func (mm *messageManager) Status() int {
	return mm.status
}

func (mm *messageManager) StartEnding() error {
	log.Printf("[MessageManager] StartEnding")
	if mm.status >= Serve_Status_Ending {
		return nil
	}
	mm.status = Serve_Status_Ending

	if err := mm.requestChan.StartEnding(); err != nil {
		return err
	}

	if err := mm.requestChan.OfficialEnding(); err != nil {
		return err
	}

	return nil
}

func (mm *messageManager) OfficialEnding() error {
	log.Printf("[MessageManager] OfficialEnding")
	mm.status = Serve_Status_Stopped

	if err := mm.responseChan.StartEnding(); err != nil {
		return err
	}

	if err := mm.responseChan.OfficialEnding(); err != nil {
		return err
	}

	return nil
}

func (mm *messageManager) DataInChannel() chan interface{} {
	return mm.requestChan.DataInChannel()
}

func (mm *messageManager) DataOutChannel() chan interface{} {
	return mm.responseChan.DataOutChannel()
}
