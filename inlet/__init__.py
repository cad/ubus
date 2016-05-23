import abc


class BaseInlet():
    def __init__(self, config, message_handler):
        self.config = config
        self.__message_handler = message_handler

    @abc.abstractmethod
    def start(self):
        pass

    @abc.abstractmethod
    def stop(self):
        pass

    def send_message(self, message):
        self.__message_handler(message)
