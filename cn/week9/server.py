# Week 9: Concurrent Echo Server using UDP (connectionless)
import socket
import threading


def handle_client(data, addr):
    msg = data.decode()
    print(threading.current_thread().name, "- Client", addr, ":", msg)
    server.sendto(data, addr)


host = input("Enter server IP: ")
port = int(input("Enter port: "))

server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind((host, port))

print("UDP Concurrent Server is running...")

while True:
    data, addr = server.recvfrom(1024)
    thread = threading.Thread(target=handle_client, args=(data, addr), daemon=True)
    thread.start()
