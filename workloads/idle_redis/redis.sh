cd idle_redis
runc run idle_redis > /dev/null &
cd ../
sleep 20
echo "start migrator"
../migrator client  --container_id idle_redis --destination 172.31.14.127 --others_path ${PWD}/idle_redis --expected_time 1 pre_copy
