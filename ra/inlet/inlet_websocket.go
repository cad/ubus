package inlet

import (
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
	"robotics.neu.edu.tr/ra27-telemetry/ra/protocol"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

const (
	INLET_TYPE_WEBSOCKET = "Websocket"
)

var WSI_ConnectionTable = map[*graceful.Server][]*websocket.Conn{}
var WSI_ServerTable = map[*Inlet]*graceful.Server{}

func MakeWebsocketInlet(config InletConfig, preproccessors_c []middleware.Config, protocol_c protocol.Config, handler InletMessageHandler) *Inlet {
	inlet := Inlet{
		Config: config,
		Handler: handler,
		Type: INLET_TYPE_WEBSOCKET,
		Preproccessors: middleware.GetProccessors(preproccessors_c),
		Protocol: protocol.GetProtocol(protocol_c),
		handlerStart: InletWebsocket_Start,
		handlerStop: InletWebsocket_Stop,
		handlerIsRunning: InletWebsocket_IsRunning,
	}
	return &inlet
}

func InletWebsocket_Start(i *Inlet) {
	var wsHandler = func (ws *websocket.Conn) {
		WSI_ConnectionTable[WSI_ServerTable[i]] = append(WSI_ConnectionTable[WSI_ServerTable[i]], ws)

		for {
			var m []byte
			if err := websocket.Message.Receive(ws, &m); err != nil {
				log.Println(err)
				if err.Error() == "EOF" {
					break
				}
			}
			i.Handler(i, m)
		}
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", websocket.Handler(wsHandler))
		srv := &graceful.Server{
			Timeout: 10 * time.Second,

			Server: &http.Server{
				Addr: i.Config["address"],
				Handler: mux,
			},
		}
		WSI_ServerTable[i] = srv
		srv.ListenAndServe()
	}()

}

func InletWebsocket_Stop(i *Inlet) {
	if val, ok := WSI_ServerTable[i]; ok {
		val.Stop(10 * time.Second)
		if _, ok := WSI_ConnectionTable[val]; ok {
			delete(WSI_ConnectionTable, val)
		}
		delete(WSI_ServerTable, i)
	}

}

func InletWebsocket_IsRunning(i *Inlet) bool {
	return true
}
