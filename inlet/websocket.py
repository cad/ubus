from autobahn.twisted.websocket import (
    WebSocketServerProtocol, WebSocketServerFactory)
from twisted.internet import reactor
from twisted.python import log
from inlet import BaseInlet


class WebSocketProtocol(WebSocketServerProtocol):
    def onConnect(self, request):
        log.msg("Client connecting: {0}".format(request.peer))

    def onOpen(self):
        self.factory.clients[self] = True
        log.msg("WebSocket connection open.")

    def onMessage(self, payload, isBinary):
        if isBinary:
            self.factory.upstream_handler(payload)
        else:
            try:
                self.factory.upstream_handler(payload.decode('utf8'))
            except Exception as e:
                log.msg(e)

    def onClose(self, wasClean, code, reason):
        del self.factory.clients[self]
        log.msg("WebSocket connection closed: {0}".format(reason))


class WebSocketFactory(WebSocketServerFactory):

    def __init__(self, uri, handler):
        WebSocketServerFactory.__init__(self, uri)
        self.clients = {}
        self.upstream_handler = handler


class WebSocketInlet(BaseInlet):
    def start(self):
        log.msg("WebSocket inlet started...")
        self.__factory = WebSocketFactory(
            u"ws://127.0.0.1:{port}".format(port=self.config['port']),
            self.send_message
        )
        self.__factory.protocol = WebSocketProtocol
        self.__interface = reactor.listenTCP(
            self.config['port'], self.__factory)

    def stop(self):
        self.protocol.close()
        log.msg("UDP inlet stopped...")
