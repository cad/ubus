from outlet import BaseOutlet


class STDOUTOutlet(BaseOutlet):
    def start(self):
        print "STDOUT outlet started..."

    def stop(self):
        print "STDOUT outlet stopped..."

    def send_message(self, message):
        print message
