# Week 8: UDP Echo Client
import socket

host = input("Enter server IP: ")
port = int(input("Enter port: "))

client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

while True:
    msg = input("Enter message: ")

    if msg.strip().lower() == "exit":
        break

    client.sendto(msg.encode(), (host, port))
    data, addr = client.recvfrom(1024)
    print("Server Echo:", data.decode())

client.close()
