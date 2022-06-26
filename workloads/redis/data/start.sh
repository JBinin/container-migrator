redis-server --bind 0.0.0.0  &
sleep 2
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 64  > /dev/null
