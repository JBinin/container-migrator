oci-runtime-tool generate --linux-cpus="0" --mounts-add '{"destination": "/data","type": "bind","source": "data","options": ["rbind","rw"]}' --process-cwd=/data --args /bin/sh --args ./start.sh > config.json

# delete net

workload
../redis.sh 