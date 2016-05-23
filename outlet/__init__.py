import abc


class BaseOutlet():
    def __init__(self, config):
        self.config = config

    @abc.abstractmethod
    def start(self):
        pass

    @abc.abstractmethod
    def stop(self):
        pass

    @abc.abstractmethod
    def send_message(self, message):
        pass
