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
}

func (c *config) Load() error {
	log.Printf("[Config] Load")
	c.name = "Mark_V4"
	c.netWork = "tcp"
	c.host = "0.0.0.0"
	c.port = "50453"
	return nil
}

func (c *config) Reload() error {
	log.Printf("[Config] Reload")
	c.name = "Mark_V4"
	c.netWork = "tcp"
	c.host = "0.0.0.0"
	c.port = "50453"
	return nil
}