/**
* @Author: Cooper
* @Date: 2019/12/13 20:01
 */

package main

import (
	"markV4/mnet"
	"time"
)

func main() {
	server , err := mnet.NewServer()
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second*time.Duration(3))

	err = server.Stop()
	if err != nil {
		panic(err)
	}
}
