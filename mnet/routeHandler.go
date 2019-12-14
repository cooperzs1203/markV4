/**
* @Author: Cooper
* @Date: 2019/12/14 19:13
 */

package mnet

import "markV4/mface"

func newRouteHandler(routeId string, routeHandleFunc RouteHandleFunc) mface.MRouteHandler {
	rh := &routeHandler{
		id:         routeId,
		handleFunc: routeHandleFunc,
	}
	return rh
}

type routeHandler struct {
	id         string
	handleFunc RouteHandleFunc
}

func (rh *routeHandler) RouteID() string {
	return rh.id
}

func (rh *routeHandler) RouteHandleFunc() func(mface.MMessage, mface.MMessage) error {
	return rh.handleFunc
}
