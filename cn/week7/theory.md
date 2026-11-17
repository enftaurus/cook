# Week 7: Concurrent Echo Server using a Connection-Oriented Protocol (TCP)

## Aim
To build a concurrent echo server with TCP. The server sends every message back to the client that sent it, and it can serve several clients at the same time.

## Theory

### TCP (Transmission Control Protocol)
- **Connection-oriented:** the client and server set up a connection (the 3-way handshake: SYN, SYN-ACK, ACK) before any data moves.
- **Reliable:** data arrives in order with no loss or duplicates. Lost segments are retransmitted.
- **Byte stream:** data is a continuous stream of bytes, not separate messages.
- **Flow and congestion control** are built in.
- Uses the socket type `SOCK_STREAM`.

### Echo server
An echo server reads whatever the client sends and sends the same data back. It is often used to test that a network connection works.

### Iterative vs. concurrent server
| Iterative server | Concurrent server |
|---|---|
| Serves one client at a time | Serves many clients at the same time |
| Other clients wait in a queue | Each client gets its own thread or process |
| Simple, low overhead | Better response time for many clients |

In this program the main thread only **accepts** connections. Each accepted connection is given to a **new thread** (`threading.Thread`), so one slow client does not block the others.

### Socket calls used
| Server | Client |
|---|---|
| `socket()` creates a TCP socket | `socket()` |
| `setsockopt(SO_REUSEADDR)` lets the server restart on the same port right away | |
| `bind()` attaches the socket to an IP and port | |
| `listen(5)` marks it as passive, with a backlog of 5 | |
| `accept()` blocks until a client connects and returns a new socket | `connect()` performs the handshake |
| `recv()` / `sendall()` | `sendall()` / `recv()` |
| `close()` | `close()` |

### Flow of the program
```
 Client                         Server (main thread)
 socket()                       socket() -> bind() -> listen()
 connect()  ---- handshake -->  accept() --> new thread(handle_client)
 sendall(msg) ------------->    recv()
 recv()     <-------------      sendall(msg)    (echo)
 sendall("exit") ---------->    loop ends, conn.close()
 close()
```

## Changes made to the original code
1. Added `SO_REUSEADDR`, so restarting the server does not fail with "Address already in use".
2. `send()` is now `sendall()`. TCP `send()` may send only part of the data, and `sendall()` sends all of it.
3. The client skips empty input. Sending an empty string sends nothing, so the client would wait in `recv()` forever.
4. Worker threads are `daemon=True`, so Ctrl+C stops the server cleanly.
5. The server prints the client's address with each message, which shows which client sent what.

## How to run
1. Open a terminal and start the server:
   ```
   cd week7
   python3 server.py
   ```
   The server binds to `127.0.0.1` (localhost) on port `5000`.
2. Open **two or more** other terminals and start a client in each:
   ```
   python3 client.py
   ```
   It connects to the same address automatically.
3. Type messages in each client. Type `exit` to disconnect a client.
4. Press `Ctrl+C` to stop the server.

## Sample output

**Server**
```
TCP Concurrent Server is running on 127.0.0.1 5000
Connected: ('127.0.0.1', 34042)
Connected: ('127.0.0.1', 34046)
Client ('127.0.0.1', 34046) : from B
Client ('127.0.0.1', 34042) : hello
Client ('127.0.0.1', 34042) : how are you
Disconnected: ('127.0.0.1', 34046)
Disconnected: ('127.0.0.1', 34042)
```

**Client 1**
```
Enter message: hello
Server Echo: hello
Enter message: how are you
Server Echo: how are you
Enter message: exit
```

**Client 2 (connected at the same time)**
```
Enter message: from B
Server Echo: from B
Enter message: exit
```

## Result
The concurrent TCP echo server was built and tested. Several clients connected at the same time, and each got its own messages echoed back.
