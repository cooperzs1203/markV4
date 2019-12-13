/**
* @Author: Cooper
* @Date: 2019/12/13 20:56
 */

package mface

type MConfig interface {
	Load() error
	Reload() error
}
