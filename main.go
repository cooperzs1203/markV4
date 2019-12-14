/**
* @Author: Cooper
* @Date: 2019/12/13 20:01
 */

package main

import (
	"log"
	"markV4/mface"
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

	log.Println(server.AddRoute("route1" , handler))

	go RouteTest(server)

	//time.Sleep(time.Second*time.Duration(10))

	time.Sleep(time.Second*time.Duration(3))

	err = server.Stop()
	if err != nil {
		panic(err)
	}
}

func handler(request mface.MMessage, response mface.MMessage) error {
	log.Println("request : " , request)
	log.Println("response : " , response)
	return nil
}

func RouteTest(s mface.MServer) {
	n := 0
	for {
		time.Sleep(time.Millisecond*time.Duration(100))
		n++
		a := &A{
			routeId: "route1",
			value:   int64(n),
		}

		if s.RouteManager().Status() < mnet.Serve_Status_Ending {
			s.RouteManager().DataInChannel() <- a
		}
	}
}

type A struct {
	routeId string
	value int64
}

func (a *A) RouteId() string {
	return a.routeId
}
