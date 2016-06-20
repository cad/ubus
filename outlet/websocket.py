try:
    import simplejson as json
except ImportError:
    import json
from autobahn.twisted.websocket import (
    WebSocketServerProtocol, WebSocketServerFactory)
from twisted.internet import reactor
from twisted.python import log

from outlet import BaseOutlet


class WebSocketProtocol(WebSocketServerProtocol):
    def onConnect(self, request):
        log.msg("Client connecting: {0}".format(request.peer))

    def onOpen(self):
        self.factory.clients[self] = True
        log.msg("WebSocket connection open.")

    def onMessage(self, payload, isBinary):
        if isBinary:
            log.msg("Binary message received: {0} bytes".format(len(payload)))
        else:
            log.msg("Text message received: {0}".format(
                payload.decode('utf8')))

    def onClose(self, wasClean, code, reason):
        del self.factory.clients[self]
        log.msg("WebSocket connection closed: {0}".format(reason))


class WebSocketFactory(WebSocketServerFactory):

    def __init__(self, uri):
        WebSocketServerFactory.__init__(self, uri)
        self.clients = {}


class WebSocketOutlet(BaseOutlet):
    def start(self):
        self.__factory = WebSocketFactory(
            u"ws://127.0.0.1:{port}".format(port=self.config['port']))
        self.__factory.protocol = WebSocketProtocol
        self.__interface = reactor.listenTCP(
            self.config['port'], self.__factory)

    def stop(self):
        self.protocol.close()

    def send_message(self, message):
        try:
            message = json.dumps(message.data)
            for client in self.__factory.clients:
                client.sendMessage(message, isBinary=False)
        except Exception as e:
            log.msg(e)
