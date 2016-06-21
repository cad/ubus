import abc


class BaseOutlet():
    def __init__(self, config):
        self.config = config
        self.__middlewares = []
        self.__setup_middlewares()

    def __setup_middlewares(self):
        middlewares = self.config.get('middlewares')
        if middlewares:
            for middleware_class, config in middlewares:
                middleware = middleware_class(config)
                self.__middlewares.append(middleware)

    @abc.abstractmethod
    def start(self):
        pass

    @abc.abstractmethod
    def stop(self):
        pass

    def apply_middlewares(self, message):
        for middleware in self.__middlewares:
            message = middleware.apply(message)
        return message

    @abc.abstractmethod
    def send_message(self, message):
        pass
