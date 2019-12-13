/**
* @Author: Cooper
* @Date: 2019/12/13 20:49
 */

package mnet

import (
	"log"
	"markV4/mface"
)

func newRouteManager() mface.MRouteManager {
	rm := &routeManager{
		status:Serve_Status_UnStarted,
	}
	return rm
}

type routeManager struct {
	status int
}

func (rm *routeManager) SetServer(s mface.MServer) {

}

func (rm *routeManager) Load() error {
	log.Printf("[RouteManager] Load")
	rm.status = Serve_Status_Load
	return nil
}

func (rm *routeManager) Start() error {
	log.Printf("[RouteManager] Start")
	rm.status = Serve_Status_Start
	return nil
}

func (rm *routeManager) Reload() error {
	log.Printf("[RouteManager] Reload")
	rm.status = Serve_Status_Reload
	return nil
}

func (rm *routeManager) Status() int {
	return rm.status
}

func (rm *routeManager) StartEnding() error {
	log.Printf("[RouteManager] StartEnding")
	rm.status = Serve_Status_Ending
	return nil
}

func (rm *routeManager) OfficialEnding() error {
	log.Printf("[RouteManager] OfficialEnding")
	rm.status = Serve_Status_Stopped
	return nil
}

func (rm *routeManager) DataInChannel() chan interface{} {
	return nil
}

func (rm *routeManager) DataOutChannel() chan interface{} {
	return nil
}
