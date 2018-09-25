# tcprfc868
Simple  Go program that runs a server which implements the Time Protocol, as described in RFC 868 . The program should fully implement the protocol as described in the RFC, except that it only needs to work over TCP, not UDP.

The program compiles to a single executable that can be run to start a server
The program takes the following flags as parameters:
-p Port for the server to listen on. Example: -p 11037
The program logs all incoming requests to standard out
As described in the RFC, the program should send the time as a 32-bit binary integer in network byte order
