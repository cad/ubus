package middleware

import "log"
import "encoding/json"

const (
	MIDDLEWARE_TYPE_JSON_DECODER = "JSONDecode"
	MIDDLEWARE_TYPE_JSON_ENCODER = "JSONEncode"
)

func GetJSONDecoderMiddleware(c ProccessorConfig) *Middleware {
	middleware := Middleware{
		Type: MIDDLEWARE_TYPE_JSON_DECODER,
		Config: c,
		handlerProccess: JSONDMiddleware_Proccess,
	}
	return &middleware
}

func GetJSONEncoderMiddleware(c ProccessorConfig) *Middleware {
	middleware := Middleware{
		Type: MIDDLEWARE_TYPE_JSON_ENCODER,
		Config: c,
		handlerProccess: JSONEMiddleware_Proccess,
	}
	return &middleware

}


func JSONDMiddleware_Proccess(m *Middleware, rawInput interface{}) interface{} {
	input, found := rawInput.([]byte)
	if !found {
		log.Panic("expected type of input: <[]byte>")
		return nil
	}
	var data interface{}
	err := json.Unmarshal(input, &data)
	if err != nil {
		log.Println(err)
	}
	return data
}

func JSONEMiddleware_Proccess(m *Middleware, input interface{}) interface{} {
	var data interface{}
	data, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
	}
	return data
}
