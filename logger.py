from twisted.internet import reactor
from twisted.python import log


class BusLogger():
    __INLETS = []
    __OUTLETS = []
    __MIDDLEWARES = []
    __CONFIG = {}

    def __init__(self, config):
        self.__CONFIG = config
        self.__setup()

    def __setup(self):
        for inlet_class, config in self.__CONFIG['inlets']:
            inlet = inlet_class(config, self.__inlet_handler)
            self.__INLETS.append(inlet)

        for outlet_class, config in self.__CONFIG['outlets']:
            outlet = outlet_class(config)
            self.__OUTLETS.append(outlet)

        for middleware_class, config in self.__CONFIG['middlewares']:
            middleware = middleware_class(config)
            self.__MIDDLEWARES.append(middleware)
            log.msg("Middleware ready: {}".format(middleware))

    def __tear_down(self):
        pass

    def __apply_middlewares(self, message):
        for middleware in self.__MIDDLEWARES:
            message = middleware.apply(message)
        return message

    def __fan_out(self, message):
        for outlet in self.__OUTLETS:
            outlet.send_message(outlet.apply_middlewares(message))

    def __inlet_handler(self, message):
        message = self.__apply_middlewares(message)
        self.__fan_out(message)

    def run(self):
        for outlet in self.__OUTLETS:
            reactor.callFromThread(outlet.start)
        for inlet in self.__INLETS:
            reactor.callFromThread(inlet.start)
        reactor.run()


if __name__ == '__main__':
    import sys
    from inlet.udp import MulticastUDPInlet
    from inlet.pcap import PCAPInlet
    from inlet.websocket import WebSocketInlet
    from outlet.stdout import STDOUTOutlet
    from outlet.websocket import WebSocketOutlet
    from middleware.json import JSONDecodeMiddleware, JSONEncodeMiddleware
    from middleware.canethernet import CANEthernetMiddleware

    config = {
        'inlets': [
            (MulticastUDPInlet, {
                'host': '239.255.60.60',
                'port': 4876,
                'middlewares': [(CANEthernetMiddleware, {})]
            }),
            (WebSocketInlet, {
                'port': 9001,
                'middlewares': [(JSONDecodeMiddleware, {})]
            }),
            # (PCAPInlet, {
            #     'filepath': 'data/sample.pcap',
            #     'postproccessors': [(CANEthernetMiddleware, {})]
            # })
        ],
        'outlets': [
            (WebSocketOutlet, {
                'port': 9000,
                'middlewares': [(JSONEncodeMiddleware, {})]
            }),
            (STDOUTOutlet, {
                'middlewares': []
            }),

        ],
        'middlewares': [
            # (CANEthernetMiddleware, {})
        ],

    }
    logger = BusLogger(config)
    log.startLogging(sys.stdout)
    logger.run()
