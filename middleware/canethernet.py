from middleware import BaseMiddleware
from protocol import Ra27CANProtocolParser


class CANEthernetMiddleware(BaseMiddleware):
    def apply(self, message):
        message = Ra27CANProtocolParser(message).parse()
        return message
