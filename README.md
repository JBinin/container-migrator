# container-migrator

## 测量网速
```iperf
iperf -s
ierf -c ip
```

```client
./migrator client  --container_id myredis --destination 172.31.92.143 --others_path /home/ubuntu/redis --expected_time 1 pre_copy
``` 

```server
./migrator server --migrated_container_dir /home/ubuntu/target
```

```
docker export $(docker create redis) | tar -C rootfs -xf -
```

```redis
 "redis-server", "--bind", "0.0.0.0"
```