/**
* @Author: Cooper
* @Date: 2019/12/13 20:26
 */

package mface

type MRouteManager interface {
	MBaseManager

	AddRoutes([]MRouteHandler) error
	AddRoute(MRouteHandler) error
	RemoveRoute(string) error
}

type MRouteHandler interface {
	RouteID() string
	RouteHandleFunc() func(MMessage, MMessage) error
}
