/**
* @Author: Cooper
* @Date: 2019/12/13 22:22
 */

package mface

type MChannel interface {
	SetSize(size uint64)
	MBaseServe
	MBaseServeEnd
	MBaseServeDataChannel
}
