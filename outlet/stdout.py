from twisted.python import log
from twisted.internet import reactor
from outlet import BaseOutlet


class STDOUTOutlet(BaseOutlet):
    def start(self):
        log.msg("STDOUT outlet started...")

    def stop(self):
        log.msg("STDOUT outlet stopped...")

    def send_message(self, message):
        reactor.callInThread(log.msg, message)
