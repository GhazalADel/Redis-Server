to run the server, use this command:\
    ./spawn_redis_server.sh
in another terminal:\
    redis-cli -h 127.0.0.1 -p 6390
more than one command:\
echo -e "ping\nping" | redis-cli -h 127.0.0.1 -p 6390