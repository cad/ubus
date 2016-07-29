from twisted.internet.protocol import DatagramProtocol
from twisted.internet import reactor
from twisted.python import log
from inlet import BaseInlet


class MulticastUDP(DatagramProtocol):
    def set_upstream_handler(self, handler):
        self.handler = handler

    def set_multicast_group(self, group, interface_ip=None):
        self.multicast_group = group
        self.interface_ip = interface_ip

    def startProtocol(self):
        """
        Called after protocol has started listening.
        """
        self.transport.setTTL(5)

        if self.interface_ip:
            self.transport.joinGroup(self.multicast_group, self.interface_ip)
        else:
            self.transport.joinGroup(self.multicast_group)

    def datagramReceived(self, data, (host, port)):
        self.handler(data)
        # print "received %r from %s:%d" % (data, host, port)
        # self.transport.write(data, (host, port))


class MulticastUDPInlet(BaseInlet):
    def start(self):
        log.msg("UDP inlet started...")
        self.__protocol = MulticastUDP()
        self.__protocol.set_upstream_handler(self.send_message)
        if self.config.get("interface_ip"):
            self.__protocol.set_multicast_group(self.config['host'],
                                                self.config['interface_ip'])
        else:
            self.__protocol.set_multicast_group(self.config['host'])

        self.__interface = None
        self.__interface = reactor.listenMulticast(
            self.config['port'],
            self.__protocol,
            listenMultiple=True
        )

    def stop(self):
        self.__interface.stopListening()
        log.msg("UDP inlet stopped...")
