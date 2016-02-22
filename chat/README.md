# chat

Write a program called chat that can be started in two modes. 

In server mode, it should accept TCP connections on a user-specified host:port,
and broadcast all received messages to all connected clients. Prefix broadcast
messages with origin address.

In client mode, it should connect to a user-specified host:port on TCP, submit
messages from stdin, and write received messages to stdout.

