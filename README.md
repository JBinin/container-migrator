# container-migrator

## 测量网速
```iperf
iperf -s
ierf -c ip
```


```server
./migrator server --migrated_container_dir /home/ubuntu/target
```

```
docker export $(docker create redis) | tar -C rootfs -xf -
```

## redis 负载实验

```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 16  > /dev/null

2022/06/24 16:26:16 client.go:141: -----------------config.json------------------
2022/06/24 16:26:16 client.go:142: data-size(KB) :  12 	 transfer time(s):  0.594550775
2022/06/24 16:26:16 client.go:143: ----------------------------------------------
2022/06/24 16:26:24 client.go:149: --------------------rootfs--------------------
2022/06/24 16:26:24 client.go:150: data-size(KB) :  121676 	 transfer time(s):  8.218554161
2022/06/24 16:26:24 client.go:151: ----------------------------------------------
2022/06/24 16:26:24 client.go:104: -----------------------------------
2022/06/24 16:26:24 client.go:105: Disk IO :  100000  KB/s
2022/06/24 16:26:24 client.go:106: Net speed:  125000  KB/s
2022/06/24 16:26:24 client.go:107: Expect memory size:  36571.42857142857 KB
2022/06/24 16:26:24 client.go:108: -----------------------------------
2022/06/24 16:26:27 client.go:161: ----------------volume----------------------
2022/06/24 16:26:27 client.go:162: data-size(KB) :  8 	 transfer time(s):  0.482568954
2022/06/24 16:26:27 client.go:163: --------------------------------------------
2022/06/24 16:26:27 client.go:234: ---------------------dump------------------------
2022/06/24 16:26:27 client.go:235: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/24 16:26:27 client.go:236: 0.266263738 	 10188 	 0.78697056
2022/06/24 16:26:27 client.go:237: -------------------------------------------------
2022/06/24 16:26:28 client.go:250: The downtime is  1.745579986s
2022/06/24 16:26:28 client.go:258: The total migration time is  12.394595404s
2022/06/24 16:26:28 client.go:28: ---------------------PrintInfo--------------------------------------
2022/06/24 16:26:28 client.go:29: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/24 16:26:28 client.go:31: 0 	 33096 		 0.169109275 	 1.642036963
2022/06/24 16:26:28 client.go:33: --------------------------------------------------------------------
```

```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 16  > /dev/null

2022/06/24 16:24:17 client.go:141: -----------------config.json------------------
2022/06/24 16:24:17 client.go:142: data-size(KB) :  12 	 transfer time(s):  0.916409161
2022/06/24 16:24:17 client.go:143: ----------------------------------------------
2022/06/24 16:24:26 client.go:149: --------------------rootfs--------------------
2022/06/24 16:24:26 client.go:150: data-size(KB) :  121676 	 transfer time(s):  8.3263079
2022/06/24 16:24:26 client.go:151: ----------------------------------------------
2022/06/24 16:24:26 client.go:104: -----------------------------------
2022/06/24 16:24:26 client.go:105: Disk IO :  100000  KB/s
2022/06/24 16:24:26 client.go:106: Net speed:  125000  KB/s
2022/06/24 16:24:26 client.go:107: Expect memory size:  36571.42857142857 KB
2022/06/24 16:24:26 client.go:108: -----------------------------------
2022/06/24 16:24:30 client.go:161: ----------------volume----------------------
2022/06/24 16:24:30 client.go:162: data-size(KB) :  8 	 transfer time(s):  0.450457357
2022/06/24 16:24:30 client.go:163: --------------------------------------------
2022/06/24 16:24:30 client.go:234: ---------------------dump------------------------
2022/06/24 16:24:30 client.go:235: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/24 16:24:30 client.go:236: 0.177414271 	 10388 	 0.762797251
2022/06/24 16:24:30 client.go:237: -------------------------------------------------
2022/06/24 16:24:31 client.go:250: The downtime is  1.65011883s
2022/06/24 16:24:31 client.go:258: The total migration time is  14.27565262s
2022/06/24 16:24:31 client.go:28: ---------------------PrintInfo--------------------------------------
2022/06/24 16:24:31 client.go:29: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/24 16:24:31 client.go:31: 0 	 39680 		 0.178346632 	 1.908664288
2022/06/24 16:24:31 client.go:31: 1 	 11844 		 0.170542248 	 1.097632633
2022/06/24 16:24:31 client.go:33: --------------------------------------------------------------------
```

```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 32  > /dev/null

2022/06/24 16:21:27 client.go:141: -----------------config.json------------------
2022/06/24 16:21:27 client.go:142: data-size(KB) :  12 	 transfer time(s):  0.541609295
2022/06/24 16:21:27 client.go:143: ----------------------------------------------
2022/06/24 16:21:35 client.go:149: --------------------rootfs--------------------
2022/06/24 16:21:35 client.go:150: data-size(KB) :  121676 	 transfer time(s):  8.015165125
2022/06/24 16:21:35 client.go:151: ----------------------------------------------
2022/06/24 16:21:35 client.go:104: -----------------------------------
2022/06/24 16:21:35 client.go:105: Disk IO :  100000  KB/s
2022/06/24 16:21:35 client.go:106: Net speed:  125000  KB/s
2022/06/24 16:21:35 client.go:107: Expect memory size:  36571.42857142857 KB
2022/06/24 16:21:35 client.go:108: -----------------------------------
2022/06/24 16:21:40 client.go:161: ----------------volume----------------------
2022/06/24 16:21:40 client.go:162: data-size(KB) :  8 	 transfer time(s):  0.473194861
2022/06/24 16:21:40 client.go:163: --------------------------------------------
2022/06/24 16:21:40 client.go:234: ---------------------dump------------------------
2022/06/24 16:21:40 client.go:235: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/24 16:21:40 client.go:236: 0.186788158 	 13440 	 0.916080189
2022/06/24 16:21:40 client.go:237: -------------------------------------------------
2022/06/24 16:21:41 client.go:250: The downtime is  1.925241472s
2022/06/24 16:21:41 client.go:258: The total migration time is  13.999441427s
2022/06/24 16:21:41 client.go:28: ---------------------PrintInfo--------------------------------------
2022/06/24 16:21:41 client.go:29: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/24 16:21:41 client.go:31: 0 	 49168 		 0.190679911 	 2.05634366
2022/06/24 16:21:41 client.go:31: 1 	 16536 		 0.088776344 	 1.154476914
2022/06/24 16:21:41 client.go:33: --------------------------------------------------------------------
```

```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 64  > /dev/null

2022/06/24 16:17:05 client.go:141: -----------------config.json------------------
2022/06/24 16:17:05 client.go:142: data-size(KB) :  12 	 transfer time(s):  0.640994472
2022/06/24 16:17:05 client.go:143: ----------------------------------------------
2022/06/24 16:17:13 client.go:149: --------------------rootfs--------------------
2022/06/24 16:17:13 client.go:150: data-size(KB) :  121676 	 transfer time(s):  8.592248141
2022/06/24 16:17:13 client.go:151: ----------------------------------------------
2022/06/24 16:17:13 client.go:104: -----------------------------------
2022/06/24 16:17:13 client.go:105: Disk IO :  100000  KB/s
2022/06/24 16:17:13 client.go:106: Net speed:  125000  KB/s
2022/06/24 16:17:13 client.go:107: Expect memory size:  36571.42857142857 KB
2022/06/24 16:17:13 client.go:108: -----------------------------------
2022/06/24 16:17:44 client.go:161: ----------------volume----------------------
2022/06/24 16:17:44 client.go:162: data-size(KB) :  8 	 transfer time(s):  0.578345445
2022/06/24 16:17:44 client.go:163: --------------------------------------------
2022/06/24 16:17:44 client.go:234: ---------------------dump------------------------
2022/06/24 16:17:44 client.go:235: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/24 16:17:44 client.go:236: 0.445602473 	 58036 	 2.183302285
2022/06/24 16:17:44 client.go:237: -------------------------------------------------
2022/06/24 16:17:44 client.go:250: The downtime is  3.873586323s
2022/06/24 16:17:44 client.go:258: The total migration time is  40.121650488s
2022/06/24 16:17:44 client.go:28: ---------------------PrintInfo--------------------------------------
2022/06/24 16:17:44 client.go:29: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/24 16:17:44 client.go:31: 0 	 81272 		 0.228784578 	 2.761927683
2022/06/24 16:17:44 client.go:31: 1 	 92300 		 0.142775351 	 2.944281922
2022/06/24 16:17:44 client.go:31: 2 	 59156 		 0.268015886 	 2.75093847
2022/06/24 16:17:44 client.go:31: 3 	 44304 		 0.20432134 	 1.969441565
2022/06/24 16:17:44 client.go:31: 4 	 50616 		 0.203104644 	 2.123048951
2022/06/24 16:17:44 client.go:31: 5 	 42788 		 0.313713506 	 1.9993120420000001
2022/06/24 16:17:44 client.go:31: 6 	 42408 		 0.355930793 	 2.129393281
2022/06/24 16:17:44 client.go:31: 7 	 46904 		 0.329610993 	 2.214310514
2022/06/24 16:17:44 client.go:31: 8 	 73168 		 0.270176618 	 2.837521194
2022/06/24 16:17:44 client.go:31: 9 	 58856 		 0.306836433 	 2.618568893
2022/06/24 16:17:44 client.go:33: --------------------------------------------------------------------
```
```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 128  > /dev/null

2022/06/24 16:11:04 client.go:141: -----------------config.json------------------
2022/06/24 16:11:04 client.go:142: data-size(KB) :  12 	 transfer time(s):  0.849836832
2022/06/24 16:11:04 client.go:143: ----------------------------------------------
2022/06/24 16:11:13 client.go:149: --------------------rootfs--------------------
2022/06/24 16:11:13 client.go:150: data-size(KB) :  121676 	 transfer time(s):  8.872182529
2022/06/24 16:11:13 client.go:151: ----------------------------------------------
2022/06/24 16:11:13 client.go:104: -----------------------------------
2022/06/24 16:11:13 client.go:105: Disk IO :  100000  KB/s
2022/06/24 16:11:13 client.go:106: Net speed:  125000  KB/s
2022/06/24 16:11:13 client.go:107: Expect memory size:  36571.42857142857 KB
2022/06/24 16:11:13 client.go:108: -----------------------------------
2022/06/24 16:12:35 client.go:161: ----------------volume----------------------
2022/06/24 16:12:35 client.go:162: data-size(KB) :  191912 	 transfer time(s):  2.929586815
2022/06/24 16:12:35 client.go:163: --------------------------------------------
2022/06/24 16:12:35 client.go:234: ---------------------dump------------------------
2022/06/24 16:12:35 client.go:235: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/24 16:12:35 client.go:236: 1.240730845 	 389876 	 7.999864142
2022/06/24 16:12:35 client.go:237: -------------------------------------------------
2022/06/24 16:12:38 client.go:250: The downtime is  15.317830837s
2022/06/24 16:12:38 client.go:258: The total migration time is  1m34.7746038s
2022/06/24 16:12:38 client.go:28: ---------------------PrintInfo--------------------------------------
2022/06/24 16:12:38 client.go:29: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/24 16:12:38 client.go:31: 0 	 173000 		 0.234288998 	 4.323033763
2022/06/24 16:12:38 client.go:31: 1 	 100952 		 0.344862609 	 3.434792983
2022/06/24 16:12:38 client.go:31: 2 	 85328 		     0.406258923 	 3.30873013
2022/06/24 16:12:38 client.go:31: 3 	 125676 		 0.394847187 	 4.027270468
2022/06/24 16:12:38 client.go:31: 4 	 106552 		 0.461919357 	 3.996606734
2022/06/24 16:12:38 client.go:31: 5 	 197496 		 0.428392625 	 5.884052995
2022/06/24 16:12:38 client.go:31: 6 	 171548 		 0.540491888 	 6.054637722
2022/06/24 16:12:38 client.go:31: 7 	 548428 		 2.197745984 	 14.403749975
2022/06/24 16:12:38 client.go:31: 8 	 331316 		 0.677433788 	 9.129160868
2022/06/24 16:12:38 client.go:31: 9 	 265884 		 0.920381607 	 8.488999833
2022/06/24 16:12:38 client.go:33: --------------------------------------------------------------------
```

```
2022/06/24 16:12:35 server.go:69: restore:myredis
2022/06/24 16:12:38 server.go:91: Restore time is  3.143806963s
2022/06/24 16:12:38 server.go:94: Handle finished.
2022/06/24 16:17:04 server.go:52: myredis
2022/06/24 16:17:44 server.go:69: restore:myredis
2022/06/24 16:17:44 server.go:91: Restore time is  663.077005ms
2022/06/24 16:17:44 server.go:94: Handle finished.
2022/06/24 16:21:27 server.go:52: myredis
2022/06/24 16:21:40 server.go:69: restore:myredis
2022/06/24 16:21:41 server.go:91: Restore time is  346.069833ms
2022/06/24 16:21:41 server.go:94: Handle finished.
2022/06/24 16:24:16 server.go:52: myredis
2022/06/24 16:24:30 server.go:69: restore:myredis
2022/06/24 16:24:31 server.go:91: Restore time is  256.109757ms
2022/06/24 16:24:31 server.go:94: Handle finished.
2022/06/24 16:26:15 server.go:52: myredis
2022/06/24 16:26:27 server.go:69: restore:myredis
2022/06/24 16:26:28 server.go:91: Restore time is  206.324453ms
2022/06/24 16:26:28 server.go:94: Handle finished.
```