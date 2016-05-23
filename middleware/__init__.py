import abc


class BaseMiddleware():
    def __init__(self, config):
        self.config = config

    @abc.abstractmethod
    def apply(self, message):
        pass
