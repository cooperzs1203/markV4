/**
* @Author: Cooper
* @Date: 2019/12/13 20:11
 */

package mface

type MServer interface {
	MBaseServe
	MBaseServeStop

	Config() MConfig
	ConnManager() MConnManager
	MessageManager() MMessageManager
	RouteManager() MRouteManager

	AddRoutes(map[string]func(MMessage, MMessage) error) error
	AddRoute(string, func(MMessage, MMessage) error) error
	RemoveRoute(string) error
}
