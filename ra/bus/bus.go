package bus

import (
	"log"
)

func MakeBus(messageHandler BusMessageHandler) *Bus {
	bus := Bus{
		bus: make(chan BusMessage),
		handlerOnMessage: messageHandler,
		stopper: make(BusStopper),
	}
	return &bus
}


func (b *Bus) Work() {
	go func() {
		var message BusMessage
		for {
			select {
			case message = <- b.bus:
				b.handlerOnMessage(message)
			case <-b.stopper:
				return
			}
		}

	}()
}

func (b *Bus) SendMessage(message BusMessage) {
	b.bus <- message
}

func (b *Bus) Shutdown(){
	log.Println("Shutting down the bus")
	b.stopper <- true
}
