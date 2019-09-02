#! /usr/bin/env python

import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('drone.firefly.com', 4399))
print(s.getsockname())