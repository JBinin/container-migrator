cd redis
runc run testredis > /dev/null &
cd ../
sleep 30
echo "start test"
../migrator predump-only  --container_id testredis  --checkpoint_path ${PWD}/image/testredis
