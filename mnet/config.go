/**
* @Author: Cooper
* @Date: 2019/12/13 20:56
 */

package mnet

import (
	"log"
	"markV4/mface"
)

func defaultConfig() mface.MConfig {
	c := &config{}
	return c
}

type config struct {
	name string
	netWork string
	host string
	port string

	cmReqCS uint64
	cmRspCS uint64
	mmReqCS uint64
	mmRspCS uint64
	rmReqCS uint64
	rmRspCS uint64
}

func (c *config) Load() error {
	log.Printf("[Config] Load")
	c.name = "Mark_V4"
	c.netWork = "tcp"
	c.host = "0.0.0.0"
	c.port = "50453"
	c.cmReqCS = 1000
	c.cmRspCS = 1000
	c.mmReqCS = 1000
	c.mmRspCS = 1000
	c.rmReqCS = 1000
	c.rmRspCS = 1000
	return nil
}

func (c *config) Reload() error {
	log.Printf("[Config] Reload")
	c.name = "Mark_V4"
	c.netWork = "tcp"
	c.host = "0.0.0.0"
	c.port = "50453"
	c.cmReqCS = 1000
	c.cmRspCS = 1000
	c.mmReqCS = 1000
	c.mmRspCS = 1000
	c.rmReqCS = 1000
	c.rmRspCS = 1000
	return nil
}

func (c *config) CMReqCS() uint64 {
	return c.cmReqCS
}

func (c *config) CMRspCS() uint64 {
	return c.cmRspCS
}

func (c *config) MMReqCS() uint64 {
	return c.mmReqCS
}

func (c *config) MMRspCS() uint64 {
	return c.mmRspCS
}

func (c *config) RMReqCS() uint64 {
	return c.rmReqCS
}

func (c *config) RMRspCS() uint64 {
	return c.rmRspCS
}