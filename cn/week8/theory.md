# Week 8: Iterative Echo Server using a Connectionless Protocol (UDP)

## Aim
To build an iterative echo server with UDP. The server handles one datagram at a time and sends it back to the sender.

## Theory

### UDP (User Datagram Protocol)
- **Connectionless:** there is no handshake. Each message (datagram) is sent on its own, with the destination address attached.
- **Unreliable:** datagrams can be lost, duplicated or arrive out of order. There are no acknowledgements or retransmissions.
- **Message-oriented:** message boundaries are kept. One `sendto()` matches one `recvfrom()`.
- Small header (8 bytes), so overhead is low. Used for DNS, DHCP, streaming, games and VoIP.
- Uses the socket type `SOCK_DGRAM`.

### TCP vs. UDP
| TCP | UDP |
|---|---|
| Connection-oriented | Connectionless |
| Reliable, ordered | Unreliable, no ordering |
| Byte stream | Datagrams (message boundaries kept) |
| `listen()`, `accept()`, `connect()` | Not needed |
| `send()` / `recv()` | `sendto()` / `recvfrom()` |
| Slower, more overhead | Faster, less overhead |

### Iterative server
An **iterative** server handles requests **one after another** in a single loop: it receives a datagram, processes it, replies, then waits for the next one. There are no threads. UDP has no connections to keep track of, so one socket can still serve many clients in turn. Each request is short, so this works well for simple services like echo.

### Socket calls used
| Server | Client |
|---|---|
| `socket(AF_INET, SOCK_DGRAM)` | `socket(AF_INET, SOCK_DGRAM)` |
| `bind()` | (the OS picks a port on the first `sendto()`) |
| `recvfrom()` returns the data and the sender's address | `sendto(msg, (host, port))` |
| `sendto(data, addr)` replies to that address | `recvfrom()` |
| | `close()` |

### Flow of the program
```
 Client                          Server
 socket()                        socket() -> bind()
 sendto(msg, server) -------->   recvfrom()  -> (data, addr)
 recvfrom()  <---------------    sendto(data, addr)   (echo)
 ... repeat ...                  ... loop handles next datagram ...
 "exit" -> close()  (nothing is sent: UDP has no connection to close)
```

## Changes made to the original code
The original code was correct. Small improvements:
1. Added `SO_REUSEADDR` so the server can restart quickly on the same port.
2. The server prints the client's address, which shows that one server handles several clients.
3. `exit` is checked after `.strip()`, so `" exit "` also works.

## How to run
1. Start the server:
   ```
   cd week8
   python3 server.py
   ```
   The server binds to `127.0.0.1` (localhost) on port `6000`.
2. In another terminal, start the client:
   ```
   python3 client.py
   ```
   It connects to the same address automatically.
3. Type messages. Type `exit` to quit the client.
4. Press `Ctrl+C` to stop the server.

> Note: if the server is not running, the UDP client gets no reply and waits in `recvfrom()`, because UDP has no connection to fail.

## Sample output

**Server**
```
UDP Iterative Server is running on 127.0.0.1 6000
Client ('127.0.0.1', 59030) : hello
Client ('127.0.0.1', 59030) : how are you
Client ('127.0.0.1', 36496) : client two here
```

**Client 1**
```
Enter message: hello
Server Echo: hello
Enter message: how are you
Server Echo: how are you
Enter message: exit
```

**Client 2**
```
Enter message: client two here
Server Echo: client two here
Enter message: exit
```

## Result
The iterative UDP echo server was built and tested. It received datagrams one by one and echoed each one back to its sender.
