# Redis-Server
It is a simple Implementation of Redis Server in Golang which capable
of handle multiple clients at the same time and handle basic
reids commands.

## Run 
Clone Project using following command
```bash
 git clone https://github.com/GhazalADel/Redis-Server.git
```
For running server enter following command
```bash
 ./redis_server.sh
```
To run a new client in another tab enter
```bash
 redis-cli -h 127.0.0.1 -p 6390
```

!!!!!!!!!!!!!!Don't forget to remove .idea folder!!!!!!!!!!!!!!!!!!!!!!!

### Usage

SET Command
```bash
 SET hi "bye"
 SET apple fruit
 SET three 3
```
GET command
```bash
 GET hi
```