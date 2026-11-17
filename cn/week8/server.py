# Week 8: Iterative Echo Server using UDP (connectionless)
import socket

host = "127.0.0.1"  # localhost
port = 6000

server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind((host, port))

print("UDP Iterative Server is running on", host, port)

while True:
    data, addr = server.recvfrom(1024)
    print("Client", addr, ":", data.decode())
    server.sendto(data, addr)
