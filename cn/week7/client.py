# Week 7: TCP Echo Client
import socket

host = "127.0.0.1"  # localhost
port = 5000

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect((host, port))

while True:
    msg = input("Enter message: ")
    if not msg:
        continue  # sending an empty message would block on recv()

    client.sendall(msg.encode())

    if msg.strip().lower() == "exit":
        break

    data = client.recv(1024).decode()
    print("Server Echo:", data)

client.close()
