# Week 9: Concurrent Echo Server using a Connectionless Protocol (UDP)

## Aim
To build a concurrent echo server with UDP. Each incoming datagram is handled in its own thread.

## Theory

### UDP recap
UDP is **connectionless**, **unreliable** and **message-oriented** (socket type `SOCK_DGRAM`). There is no `listen()`, `accept()` or `connect()`. The server uses `recvfrom()` to get a datagram and the sender's address, and `sendto()` to reply.

### Concurrent UDP server
- In an **iterative** UDP server (Week 8), the main loop receives a datagram, handles it and replies before it reads the next one. If handling one request takes a long time, every other client waits.
- In a **concurrent** UDP server, the main loop only **receives** datagrams. Each datagram is passed to a **new thread** that handles it and sends the reply. The main loop goes straight back to `recvfrom()`, so many requests are handled at the same time.
- UDP has no per-client connection socket (unlike TCP's `accept()`), so **all threads share the one server socket** and reply with `server.sendto(data, addr)`. Sending on a shared UDP socket from several threads is safe, because each `sendto()` is a separate datagram.

### Iterative vs. concurrent (UDP)
| Iterative UDP (Week 8) | Concurrent UDP (Week 9) |
|---|---|
| Single loop, no threads | New thread per datagram |
| Requests handled one by one | Requests handled at the same time |
| Simple, fine for quick requests | Better when a request takes time to process |

### TCP concurrent vs. UDP concurrent
| TCP concurrent (Week 7) | UDP concurrent (Week 9) |
|---|---|
| One thread per **client connection** | One thread per **datagram** |
| Each thread has its own `conn` socket | All threads share one socket |
| Thread lives until the client disconnects | Thread ends after one reply |

### Flow of the program
```
 Client                          Server (main thread)
 socket()                        socket() -> bind()
 sendto(msg, server) -------->   recvfrom() -> (data, addr)
                                     |
                                     +--> new thread: handle_client(data, addr)
 recvfrom()  <---------------             server.sendto(data, addr)   (echo)
 ... repeat ...                  main thread goes back to recvfrom()
 "exit" -> close()
```

## Changes made to the original code
The original code was correct. Small improvements:
1. Added `SO_REUSEADDR` so the server can restart quickly on the same port.
2. The server prints the **thread name** with each message. This shows that a different thread handles each datagram, which is what makes it concurrent.
3. Threads are `daemon=True`, so Ctrl+C stops the server cleanly.
4. `exit` is checked after `.strip()`.

## How to run
1. Start the server:
   ```
   cd week9
   python3 server.py
   ```
   Enter `127.0.0.1` and a port such as `7000`.
2. Open **two or more** terminals and start a client in each:
   ```
   python3 client.py
   ```
   Enter the same IP and port.
3. Type messages in each client. Type `exit` to quit a client.
4. Press `Ctrl+C` to stop the server.

## Sample output

**Server**
```
Enter server IP: 127.0.0.1
Enter port: 7000
UDP Concurrent Server is running...
Thread-1 (handle_client) - Client ('127.0.0.1', 54095) : hello
Thread-2 (handle_client) - Client ('127.0.0.1', 54095) : how are you
Thread-3 (handle_client) - Client ('127.0.0.1', 48110) : client two here
```

**Client 1**
```
Enter server IP: 127.0.0.1
Enter port: 7000
Enter message: hello
Server Echo: hello
Enter message: how are you
Server Echo: how are you
Enter message: exit
```

**Client 2**
```
Enter server IP: 127.0.0.1
Enter port: 7000
Enter message: client two here
Server Echo: client two here
Enter message: exit
```

## Result
The concurrent UDP echo server was built and tested. Each datagram was handled by its own thread and echoed back to the right client.
