import dpkt
import protocol
import sys, traceback

f = open('data/sample.pcap')
pcap = dpkt.pcap.Reader(f)

for ts, buf in pcap:
    eth = dpkt.ethernet.Ethernet(buf)
    ip = eth.data
    udp = ip.data
    try:
        obj = protocol.Ra27CANProtocolParser(udp.data).parse()
        print (obj)
    except Exception as e:
        traceback.print_exc(file=sys.stdout)
        print "Exception", e

f.close()
