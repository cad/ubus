from inlet.udp import MulticastUDPInlet
from inlet.pcap import PCAPInlet
from inlet.websocket import WebSocketInlet
from outlet.stdout import STDOUTOutlet
from outlet.websocket import WebSocketOutlet
from middleware.json import JSONDecodeMiddleware, JSONEncodeMiddleware
from middleware.canethernet import CANEthernetMiddleware

CONFIG = {
    'inlets': [
        (MulticastUDPInlet, {
            'host': '239.255.60.60',
            'port': 4876,
            'middlewares': [(CANEthernetMiddleware, {})]
        }),
        (WebSocketInlet, {
            'port': 9001,
            'middlewares': [(JSONDecodeMiddleware, {})]
        }),
        # (PCAPInlet, {
        #     'filepath': 'data/sample.pcap',
        #     'middlewares': [(CANEthernetMiddleware, {})]
        # })
    ],
    'outlets': [
        (WebSocketOutlet, {
            'port': 9000,
            'middlewares': [(JSONEncodeMiddleware, {})]
        }),
        (STDOUTOutlet, {
            'middlewares': []
        }),

    ],
    'middlewares': [
        # (CANEthernetMiddleware, {})
    ],

}
