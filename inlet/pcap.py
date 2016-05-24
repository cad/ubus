import dpkt
import traceback
import sys
from twisted.python import log
from inlet import BaseInlet


class PCAPInlet(BaseInlet):
    def start(self):
        log.msg("PCAP inlet started...")
        self.__pcap_file = open(self.config['filepath'])
        self.__pcap_reader = dpkt.pcap.Reader(self.__pcap_file)

        for ts, buf in self.__pcap_reader:
            eth = dpkt.ethernet.Ethernet(buf)
            ip = eth.data
            udp = ip.data
            try:
                udp.data
                self.send_message(udp.data)
            except Exception as e:
                traceback.print_exc(file=sys.stdout)
                print "Exception", e

        self.__pcap_file.close()

    def stop(self):
        self.__interface.stopListening()
        log.msg("PCAP inlet stopped...")
