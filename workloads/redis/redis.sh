cd redis
runc run myredis > /dev/null &
cd ../
sleep 10
../migrator client  --container_id myredis --destination 172.31.14.127 --others_path ${PWD} --expected_time 1 pre_copy