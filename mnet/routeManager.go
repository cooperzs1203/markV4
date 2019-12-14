/**
* @Author: Cooper
* @Date: 2019/12/13 20:49
 */

package mnet

import (
	"errors"
	"fmt"
	"log"
	"markV4/mface"
	"sync"
)

type RouteHandleFunc func(mface.MMessage, mface.MMessage) error

func newRouteManager() mface.MRouteManager {
	rm := &routeManager{
		status:       Serve_Status_UnStarted,
		server:       nil,
		requestChan:  newChannel(),
		responseChan: newChannel(),
		routes:       make(map[string]mface.MRouteHandler),
		routeMu:      sync.RWMutex{},
	}
	return rm
}

type routeManager struct {
	status int
	server mface.MServer

	requestChan  mface.MChannel
	responseChan mface.MChannel

	routes  map[string]mface.MRouteHandler
	routeMu sync.RWMutex
}

func (rm *routeManager) SetServer(s mface.MServer) {
	rm.server = s
}

func (rm *routeManager) Load() error {
	log.Printf("[RouteManager] Load")
	rm.status = Serve_Status_Load

	rm.requestChan.SetSize(rm.server.Config().RMReqCS())
	if err := rm.requestChan.Load(); err != nil {
		return err
	}

	rm.responseChan.SetSize(rm.server.Config().RMRspCS())
	if err := rm.responseChan.Load(); err != nil {
		return err
	}

	return nil
}

func (rm *routeManager) Start() error {
	log.Printf("[RouteManager] Start")
	rm.status = Serve_Status_Start

	if err := rm.requestChan.Start(); err != nil {
		return err
	}

	if err := rm.responseChan.Start(); err != nil {
		return err
	}

	go rm.startAcceptRequest()

	return nil
}

func (rm *routeManager) Reload() error {
	log.Printf("[RouteManager] Reload")
	rm.status = Serve_Status_Reload

	rm.requestChan.SetSize(rm.server.Config().RMReqCS())
	if err := rm.requestChan.Reload(); err != nil {
		return err
	}

	rm.responseChan.SetSize(rm.server.Config().RMRspCS())
	if err := rm.responseChan.Reload(); err != nil {
		return err
	}

	return nil
}

func (rm *routeManager) Status() int {
	return rm.status
}

func (rm *routeManager) StartEnding() error {
	log.Printf("[RouteManager] StartEnding")
	if rm.status >= Serve_Status_Ending {
		return nil
	}
	rm.status = Serve_Status_Ending

	if err := rm.requestChan.StartEnding(); err != nil {
		return err
	}

	if err := rm.requestChan.OfficialEnding(); err != nil {
		return err
	}
	return nil
}

func (rm *routeManager) OfficialEnding() error {
	log.Printf("[RouteManager] OfficialEnding")
	rm.status = Serve_Status_Stopped

	if err := rm.responseChan.StartEnding(); err != nil {
		return err
	}

	if err := rm.responseChan.OfficialEnding(); err != nil {
		return err
	}

	return nil
}

func (rm *routeManager) DataInChannel() chan interface{} {
	return rm.requestChan.DataInChannel()
}

func (rm *routeManager) DataOutChannel() chan interface{} {
	return rm.responseChan.DataOutChannel()
}

func (rm *routeManager) AddRoutes(routes []mface.MRouteHandler) error {
	rm.routeMu.Lock()
	defer rm.routeMu.Unlock()

	for index := range routes {
		route := routes[index]
		if _, ok := rm.routes[route.RouteID()]; ok {
			return errors.New(fmt.Sprintf("[%s] route exists", route.RouteID()))
		}
		rm.routes[route.RouteID()] = route
	}

	return nil
}

func (rm *routeManager) AddRoute(route mface.MRouteHandler) error {
	rm.routeMu.Lock()
	defer rm.routeMu.Unlock()

	if _, ok := rm.routes[route.RouteID()]; ok {
		return errors.New(fmt.Sprintf("[%s] route exists", route.RouteID()))
	}

	rm.routes[route.RouteID()] = route

	return nil
}

func (rm *routeManager) RemoveRoute(routeId string) error {
	rm.routeMu.Lock()
	defer rm.routeMu.Unlock()

	delete(rm.routes, routeId)

	return nil
}

// ============= private methods mark

func (rm *routeManager) startAcceptRequest() {
	for {
		if !rm.requestChan.DataOutChannelWorking() {
			continue
		}
		request , ok := <- rm.requestChan.DataOutChannel()
		if !ok {
			if rm.status >= Serve_Status_Ending {
				break
			} else {
				continue
			}
		}

		if request, ok := request.(mface.MMessage); ok {
			rm.handleRequest(request)
		}
	}
}

func (rm *routeManager) handleRequest(request mface.MMessage) {
	if request == nil || request.RouteId() == "" {
		return
	}

	route , err := rm.handleRoute(request.RouteId())
	if err != nil {
		log.Printf("[RouteManager] get %s route handler error : %+v" , request.RouteId() , err)
		return
	}

	go rm.execRouteHandleFunc(request , route)
}

func (rm *routeManager) handleRoute(routeId string) (mface.MRouteHandler, error) {
	rm.routeMu.RLock()
	defer rm.routeMu.RUnlock()

	var err error
	route, exists := rm.routes[routeId]
	if !exists {
		err = errors.New(fmt.Sprintf("[%s] routeId not exists", routeId))
	}

	return route, err
}

func (rm *routeManager) execRouteHandleFunc(request mface.MMessage, route mface.MRouteHandler) {
	var response mface.MMessage
	handleFunc := route.RouteHandleFunc()
	err := handleFunc(request , response)
	if err != nil {
		// todo:RouteHandleFunc 的 error 返回值需要重做，符合mface.MMessage
		return
	}
	if rm.responseChan.DataInChannelWorking() {
		rm.responseChan.DataInChannel() <- response
	}
}