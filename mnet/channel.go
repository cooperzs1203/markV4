/**
* @Author: Cooper
* @Date: 2019/12/13 21:27
 */

package mnet

import (
	"log"
)

func newChannel() *channel {
	c := &channel{
		status:      Serve_Status_UnStarted,
		size:        0,
		dataInChan:  nil,
		dataOutChan: nil,
	}
	return c
}

type channel struct {
	status      int
	size        uint64
	dataInChan  *chan interface{}
	dataOutChan *chan interface{}
}

func (c *channel) SetSize(size uint64) {
	c.size = size
}

func (c *channel) Load() error {
	log.Printf("[Channel] Load")
	c.status = Serve_Status_Load

	dataInChan := make(chan interface{}, c.size)
	dataOutChan := make(chan interface{}, c.size)
	c.dataInChan = &dataInChan
	c.dataOutChan = &dataOutChan

	return nil
}

func (c *channel) Start() error {
	log.Printf("[Channel] Start")
	c.status = Serve_Status_Start

	go c.start()

	return nil
}

func (c *channel) Reload() error {
	log.Printf("[Channel] Reload")
	c.status = Serve_Status_Reload

	newDataInChan := make(chan interface{}, c.size)
	newDataOutChan := make(chan interface{}, c.size)
	oldDataInChan := c.dataInChan
	oldDataOutChan := c.dataOutChan
	c.dataInChan = &newDataInChan
	c.dataOutChan = &newDataOutChan

	exportCache := func(newChan, oldChan *chan interface{}, finish chan bool) {
		close(*oldChan)
		for {
			data, ok := <-*oldChan
			if !ok {
				break
			}
			*newChan <- data
		}
		finish <- true
	}

	finishChan := make(chan bool, 0)
	go exportCache(c.dataInChan, oldDataInChan, finishChan)
	go exportCache(c.dataOutChan, oldDataOutChan, finishChan)

	signalNumber := 0
	for {
		<-finishChan
		signalNumber++
		if signalNumber == 2 {
			break
		}
	}

	c.status = Serve_Status_Start

	return nil
}

func (c *channel) Status() int {
	return c.status
}

func (c *channel) StartEnding() error {
	if c.status == Serve_Status_Reload {
		for {
			if c.status == Serve_Status_Start {
				break
			}
		}
	}
	log.Printf("[Channel] StartEnding")
	if c.status >= Serve_Status_Ending {
		return nil
	}
	c.status = Serve_Status_Ending
	close(*c.dataInChan)

	for {
		if len(*c.dataInChan) == 0 {
			break
		}
	}

	c.dataInChan = nil

	return nil
}

func (c *channel) OfficialEnding() error {
	log.Printf("[Channel] OfficialEnding")
	c.status = Serve_Status_Stopped

	close(*c.dataOutChan)

	for {
		if len(*c.dataOutChan) == 0 {
			break
		}
	}

	return nil
}

func (c *channel) DataInChannel() chan interface{} {
	return *c.dataInChan
}

func (c *channel) DataOutChannel() chan interface{} {
	return *c.dataOutChan
}

func (c *channel) start() {
	for {
		data, ok := <-*c.dataInChan
		if !ok {
			if c.status >= Serve_Status_Ending {
				break
			} else {
				continue
			}
		}
		*c.dataOutChan <- data
	}
}
