# Week 7: Concurrent Echo Server using TCP (connection-oriented)
import socket
import threading


def handle_client(conn, addr):
    print("Connected:", addr)
    while True:
        data = conn.recv(1024).decode()
        if not data or data.strip().lower() == "exit":
            break
        print("Client", addr, ":", data)
        conn.sendall(data.encode())
    conn.close()
    print("Disconnected:", addr)


host = input("Enter server IP: ")
port = int(input("Enter port: "))

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
server.bind((host, port))
server.listen(5)

print("TCP Concurrent Server is running...")

while True:
    conn, addr = server.accept()
    thread = threading.Thread(target=handle_client, args=(conn, addr), daemon=True)
    thread.start()
