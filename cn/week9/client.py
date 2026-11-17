# Week 9: UDP Echo Client
import socket

host = "127.0.0.1"  # localhost
port = 7000

client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

while True:
    msg = input("Enter message: ")

    if msg.strip().lower() == "exit":
        break

    client.sendto(msg.encode(), (host, port))
    data, addr = client.recvfrom(1024)
    print("Server Echo:", data.decode())

client.close()
