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
		status:Serve_Status_UnStarted,
	}
	return cm
}

type connManager struct {
	status int
}

func (cm *connManager) SetServer(s mface.MServer) {

}

func (cm *connManager) Load() error {
	log.Printf("[ConnManager] Load")
	cm.status = Serve_Status_Load
	return nil
}

func (cm *connManager) Start() error {
	log.Printf("[ConnManager] Start")
	cm.status = Serve_Status_Start
	return nil
}

func (cm *connManager) Reload() error {
	log.Printf("[ConnManager] Reload")
	cm.status = Serve_Status_Reload
	return nil
}

func (cm *connManager) Status() int {
	return cm.status
}

func (cm *connManager) StartEnding() error {
	log.Printf("[ConnManager] StartEnding")
	cm.status = Serve_Status_Ending
	return nil
}

func (cm *connManager) OfficialEnding() error {
	log.Printf("[ConnManager] OfficialEnding")
	cm.status = Serve_Status_Stopped
	return nil
}

func (cm *connManager) DataInChannel() chan interface{} {
	return nil
}

func (cm *connManager) DataOutChannel() chan interface{} {
	return nil
}