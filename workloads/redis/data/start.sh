redis-server --bind 0.0.0.0  &
sleep 2
redis-benchmark -t SET -c 20 -n 3000000 -r 10000000 -d 2048  > /dev/null
