/**
* @Author: Cooper
* @Date: 2019/12/13 20:17
 */

package mnet

import (
	"log"
	"markV4/mface"
)

func newConnManager() mface.MConnManager {
	cm := &connManager{
		status:       Serve_Status_UnStarted,
		server:       nil,
		requestChan:  newChannel(),
		responseChan: newChannel(),
	}
	return cm
}

type connManager struct {
	status int
	server mface.MServer

	requestChan  mface.MChannel
	responseChan mface.MChannel
}

func (cm *connManager) SetServer(s mface.MServer) {
	cm.server = s
}

func (cm *connManager) Load() error {
	log.Printf("[ConnManager] Load")
	cm.status = Serve_Status_Load

	cm.requestChan.SetSize(cm.server.Config().CMReqCS())
	if err := cm.requestChan.Load(); err != nil {
		return err
	}

	cm.responseChan.SetSize(cm.server.Config().CMRspCS())
	if err := cm.responseChan.Load(); err != nil {
		return err
	}

	return nil
}

func (cm *connManager) Start() error {
	log.Printf("[ConnManager] Start")
	cm.status = Serve_Status_Start

	if err := cm.requestChan.Start(); err != nil {
		return err
	}

	if err := cm.responseChan.Start(); err != nil {
		return err
	}

	return nil
}

func (cm *connManager) Reload() error {
	log.Printf("[ConnManager] Reload")
	cm.status = Serve_Status_Reload

	cm.requestChan.SetSize(cm.server.Config().CMReqCS())
	if err := cm.requestChan.Reload(); err != nil {
		return err
	}

	cm.responseChan.SetSize(cm.server.Config().CMRspCS())
	if err := cm.responseChan.Reload(); err != nil {
		return err
	}

	return nil
}

func (cm *connManager) Status() int {
	return cm.status
}

func (cm *connManager) StartEnding() error {
	log.Printf("[ConnManager] StartEnding")
	if cm.status >= Serve_Status_Ending {
		return nil
	}
	cm.status = Serve_Status_Ending

	if err := cm.requestChan.StartEnding(); err != nil {
		return err
	}

	if err := cm.requestChan.OfficialEnding(); err != nil {
		return err
	}

	return nil
}

func (cm *connManager) OfficialEnding() error {
	log.Printf("[ConnManager] OfficialEnding")
	cm.status = Serve_Status_Stopped

	if err := cm.responseChan.StartEnding(); err != nil {
		return err
	}

	if err := cm.responseChan.OfficialEnding(); err != nil {
		return err
	}

	return nil
}

func (cm *connManager) DataInChannel() chan interface{} {
	return cm.requestChan.DataInChannel()
}

func (cm *connManager) DataOutChannel() chan interface{} {
	return cm.responseChan.DataOutChannel()
}
