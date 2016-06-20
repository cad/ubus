import dpkt
import traceback
import sys
from twisted.python import log
from twisted.internet import reactor
from inlet import BaseInlet


class PCAPInlet(BaseInlet):
    def start(self):
        log.msg("PCAP inlet started...")

        def get_packets():
            self.__pcap_file = open(self.config['filepath'])
            self.__pcap_reader = dpkt.pcap.Reader(self.__pcap_file)
            pcap = list(self.__pcap_reader)
            self.__pcap_file.close()
            return pcap

        pcap = get_packets()
        old_ts = 0
        index = 0

        def proccess(old_ts, index):
            if index >= len(pcap):
                return
            ts, buf = pcap[index]
            eth = dpkt.ethernet.Ethernet(buf)
            ip = eth.data
            udp = ip.data
            try:
                dt = ts - old_ts if old_ts else 0
                self.send_message(udp.data)
                old_ts = ts
                index += 1
                reactor.callLater(dt, proccess, old_ts, index)
            except Exception as e:
                traceback.print_exc(file=sys.stdout)
                print "Exception", e

        proccess(old_ts, index)

    def stop(self):
        self.__interface.stopListening()
        log.msg("PCAP inlet stopped...")
