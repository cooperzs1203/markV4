/**
* @Author: Cooper
* @Date: 2019/12/13 20:02
 */

package mface

type MBaseServe interface {
	Load() error
	Start() error
	Reload() error
	Status() int
}

type MBaseServeStop interface {
	Stop() error
}

type MBaseServeEnd interface {
	StartEnding() error
	OfficialEnding() error
}

type MBaseServeDataChannel interface {
	DataInChannel() chan interface{}
	DataOutChannel() chan interface{}
}

type MBaseManager interface {
	MBaseServe
	MBaseServeEnd
	MBaseServeDataChannel
	SetServer(MServer)
}