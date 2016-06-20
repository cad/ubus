from middleware import BaseMiddleware
from protocol import Ra27CANProtocolParser


class CANEthernetMiddleware(BaseMiddleware):
    def apply(self, message):
        try:
            message = Ra27CANProtocolParser(message).parse()
        except Exception as e:
            print e
        return message
