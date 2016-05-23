from twisted.internet import reactor


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

    def __tear_down(self):
        pass

    def __apply_middlewares(self, message):
        for middleware in self.__MIDDLEWARES:
            message = middleware.apply(message)
        return message

    def __fan_out(self, message):
        for outlet in self.__OUTLETS:
            outlet.send_message(message)

    def __inlet_handler(self, message):
        message = self.__apply_middlewares(message)
        self.__fan_out(message)

    def run(self):
        for outlet in self.__OUTLETS:
            outlet.start()
        for inlet in self.__INLETS:
            inlet.start()
        reactor.run()




if __name__ == '__main__':
    from inlet.udp import MulticastUDPInlet
    from outlet.stdout import STDOUTOutlet

    config = {
        'inlets': [
            (MulticastUDPInlet, {
                'host': '239.255.60.60',
                'port': 4876,
            })
        ],
        'outlets': [
            (STDOUTOutlet, {

            })
        ],
        'middlewares': [
        ],

    }
    logger = BusLogger(config)
    logger.run()
