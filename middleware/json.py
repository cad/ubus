from middleware import BaseMiddleware
import simplejson as json


class JSONDecodeMiddleware(BaseMiddleware):
    def apply(self, message):
        try:
            message = json.loads(message)
        except Exception as e:
            print e
        return message


class JSONEncodeMiddleware(BaseMiddleware):
    def apply(self, message):
        try:
            message = json.dumps(message)
        except Exception as e:
            print e
        return message
