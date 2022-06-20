# container-migrator

## net
```iperf
iperf -s
ierf -c ip
```

```client
./migrator client  --container_id mybusybox --destination 172.31.92.143  --others_path /home/ubuntu/container pre_copy
```

```server
./migrator server --migrated_container_dir /home/ubuntu/target
```