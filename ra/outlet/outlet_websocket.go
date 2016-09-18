package outlet

import (
	"log"
	"time"
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"golang.org/x/net/websocket"
)

const (
	OUTLET_TYPE_WEBSOCKET = "Websocket"
)

type WebsocketOutlet Outlet
var WSO_ConnectionTable = map[*graceful.Server][]*websocket.Conn{}
var WSO_ServerTable = map[*Outlet]*graceful.Server{}

func MakeWebsocketOutlet(config OutletConfig, preproccessors_c []middleware.Config) *Outlet {
	outlet := Outlet{
		Config: config,
		Type: OUTLET_TYPE_WEBSOCKET,
		Preproccessors: middleware.GetProccessors(preproccessors_c),
		handlerStart: OutletWebsocket_Start,
		handlerStop: OutletWebsocket_Stop,
		handlerDrain: OutletWebsocket_Drain,
		handlerIsRunning: OutletWebsocket_IsRunning,
	}
	return &outlet
}

func OutletWebsocket_Start(o *Outlet) {
	var wsHandler = func (ws *websocket.Conn) {
		WSO_ConnectionTable[WSO_ServerTable[o]] = append(WSO_ConnectionTable[WSO_ServerTable[o]], ws)

		for {
			var m interface{}
			if err := websocket.JSON.Receive(ws, &m); err != nil {
				log.Println(err)
				if err.Error() == "EOF" {
					break
				}
			}
			log.Println("Received message:", m)
		}
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", websocket.Handler(wsHandler))
		srv := &graceful.Server{
			Timeout: 10 * time.Second,

			Server: &http.Server{
				Addr: o.Config["address"],
				Handler: mux,
			},
		}
		WSO_ServerTable[o] = srv
		srv.ListenAndServe()
	}()
}

func OutletWebsocket_Stop(o *Outlet) {
	if val, ok := WSO_ServerTable[o]; ok {
		val.Stop(10 * time.Second)
		if _, ok := WSO_ConnectionTable[val]; ok {
			delete(WSO_ConnectionTable, val)
		}
		delete(WSO_ServerTable, o)
	}
}

func OutletWebsocket_IsRunning(o *Outlet) bool {
	return true
}

func OutletWebsocket_Drain(o *Outlet, message OutletMessage) {
	connections := WSO_ConnectionTable[WSO_ServerTable[o]]
	for i, c := range connections  {

		if err := websocket.JSON.Send(c, message); err != nil {
			log.Println(err)

			// Remove connection from the slice
			a := WSO_ConnectionTable[WSO_ServerTable[o]]
			a[i] = a[len(a)-1]
			a = a[:len(a)-1]
			WSO_ConnectionTable[WSO_ServerTable[o]] = a
		}
	}
}
