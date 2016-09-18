package middleware

type Config struct {
	Type string `json:"type"`
	Config ProccessorConfig
}

type ProccessorConfig map[string]string

type Middleware struct {
	Type string
	Config ProccessorConfig
	handlerProccess MiddlewareMessageHandler
}

type MiddlewareMessageHandler func(*Middleware, interface{}) interface{}

type Proccessor interface {
	Proccess(interface{}) interface{}
}
