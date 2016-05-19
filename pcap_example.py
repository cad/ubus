import dpkt
import protocol

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
        print "Exception", e
f.close()
