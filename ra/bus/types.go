package bus

type Bus struct {
	bus chan BusMessage
	handlerOnMessage BusMessageHandler
	stopper BusStopper
}
type BusMessage map[string]interface{}
type BusStopper chan bool
type BusMessageHandler func(BusMessage)
