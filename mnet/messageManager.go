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
		status:Serve_Status_UnStarted,
	}
	return mm
}

type messageManager struct {
	status int
}

func (mm *messageManager) SetServer(s mface.MServer) {

}

func (mm *messageManager) Load() error {
	log.Printf("[MessageManager] Load")
	mm.status = Serve_Status_Load
	return nil
}

func (mm *messageManager) Start() error {
	log.Printf("[MessageManager] Start")
	mm.status = Serve_Status_Start
	return nil
}

func (mm *messageManager) Reload() error {
	log.Printf("[MessageManager] Reload")
	mm.status = Serve_Status_Reload
	return nil
}

func (mm *messageManager) Status() int {
	return mm.status
}

func (mm *messageManager) StartEnding() error {
	log.Printf("[MessageManager] StartEnding")
	mm.status = Serve_Status_Ending
	return nil
}

func (mm *messageManager) OfficialEnding() error {
	log.Printf("[MessageManager] OfficialEnding")
	mm.status = Serve_Status_Stopped
	return nil
}

func (mm *messageManager) DataInChannel() chan interface{} {
	return nil
}

func (mm *messageManager) DataOutChannel() chan interface{} {
	return nil
}
