/**
* @Author: Cooper
* @Date: 2019/12/13 20:56
 */

package mface

type MConfig interface {
	Load() error
	Reload() error

	CMReqCS() uint64
	CMRspCS() uint64
	MMReqCS() uint64
	MMRspCS() uint64
	RMReqCS() uint64
	RMRspCS() uint64
}
