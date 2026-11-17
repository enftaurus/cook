# Week 8: Iterative Echo Server using UDP (connectionless)
import socket

host = input("Enter server IP: ")
port = int(input("Enter port: "))

server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind((host, port))

print("UDP Iterative Server is running...")

while True:
    data, addr = server.recvfrom(1024)
    print("Client", addr, ":", data.decode())
    server.sendto(data, addr)
