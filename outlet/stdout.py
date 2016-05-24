from twisted.python import log
from outlet import BaseOutlet


class STDOUTOutlet(BaseOutlet):
    def start(self):
        log.msg("STDOUT outlet started...")

    def stop(self):
        log.msg("STDOUT outlet stopped...")

    def send_message(self, message):
        log.msg(message)
