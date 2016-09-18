package middleware

import "log"

func GetProccessor(proccessorType string, c ProccessorConfig) *Middleware {
	switch proccessorType {
	case MIDDLEWARE_TYPE_JSON_DECODER:
		return GetJSONDecoderMiddleware(c)
	case MIDDLEWARE_TYPE_JSON_ENCODER:
		return GetJSONEncoderMiddleware(c)
	default:
		log.Fatalf("middleware type not defined: <%s> ", proccessorType)
		return nil
	}
}

func GetProccessors(c []Config) *[]*Middleware{
	var proccessors []*Middleware
	for _, proccessor := range c {
		proccessors = append(proccessors, GetProccessor(proccessor.Type, proccessor.Config))
	}

	return &proccessors
}

func (m *Middleware) Proccess(input interface{}) interface{} {
	execute := m.handlerProccess
	return execute(m, input)
}
