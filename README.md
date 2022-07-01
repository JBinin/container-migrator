# container-migrator

## 测量网速
```iperf
iperf -s
ierf -c ip
```


```server
./container-migrator/migrator server --migrated_container_dir /home/ubuntu/target
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

## redis disk-less dedup
```
redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 64  > /dev/null

2022/06/28 02:12:56 client.go:149: -----------------config.json------------------
2022/06/28 02:12:56 client.go:150: data-size(KB) :  12 	 transfer time(s):  0.459447751
2022/06/28 02:12:56 client.go:151: ----------------------------------------------
2022/06/28 02:13:05 client.go:157: --------------------rootfs--------------------
2022/06/28 02:13:05 client.go:158: data-size(KB) :  121676 	 transfer time(s):  8.783336112
2022/06/28 02:13:05 client.go:159: ----------------------------------------------
2022/06/28 02:13:05 client.go:112: -----------------------------------
2022/06/28 02:13:05 client.go:113: Disk IO :  100000  KB/s
2022/06/28 02:13:05 client.go:114: Net speed:  125000  KB/s
2022/06/28 02:13:05 client.go:115: Expect memory size:  36571.42857142857 KB
2022/06/28 02:13:05 client.go:116: -----------------------------------
2022/06/28 02:13:25 client.go:169: ----------------volume----------------------
2022/06/28 02:13:25 client.go:170: data-size(KB) :  5752 	 transfer time(s):  0.474880682
2022/06/28 02:13:25 client.go:171: --------------------------------------------
2022/06/28 02:13:25 client.go:250: ---------------------dump------------------------
2022/06/28 02:13:25 client.go:251: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/06/28 02:13:25 client.go:252: 0.193347906 	 11876 	 0.700287606
2022/06/28 02:13:25 client.go:253: -------------------------------------------------
2022/06/28 02:13:25 client.go:266: The downtime is  1.650690135s
2022/06/28 02:13:25 client.go:274: The total migration time is  28.834866067s
2022/06/28 02:13:25 client.go:29: ---------------------PrintInfo--------------------------------------
2022/06/28 02:13:25 client.go:30: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/06/28 02:13:25 client.go:32: 0 	 158076 		 0.234491292 	 5.313152193
2022/06/28 02:13:25 client.go:32: 1 	 171332 		 0.226328632 	 5.626286134
2022/06/28 02:13:25 client.go:32: 2 	 134960 		 0.357991926 	 5.198836905
2022/06/28 02:13:25 client.go:32: 3 	 15980 		 0.075746584 	 0.875574756
2022/06/28 02:13:25 client.go:34: --------------------------------------------------------------------
```

## motivation

redis-benchmark -t SET -c 10 -n 10000000 -r 10000000 -d 128  > /dev/null

```
2022/07/01 12:42:48 server.go:52: myredis
2022/07/01 12:44:12 server.go:75: restore:myredis
2022/07/01 12:44:14 server.go:105: Restore time is  1.98553431s
2022/07/01 12:44:14 server.go:108: Handle finished.
```

```
2022/07/01 12:42:49 client.go:149: -----------------config.json------------------
2022/07/01 12:42:49 client.go:150: data-size(KB) :  12 	 transfer time(s):  0.615403514
2022/07/01 12:42:49 client.go:151: ----------------------------------------------
2022/07/01 12:43:00 client.go:157: --------------------rootfs--------------------
2022/07/01 12:43:00 client.go:158: data-size(KB) :  121680 	 transfer time(s):  10.807679052
2022/07/01 12:43:00 client.go:159: ----------------------------------------------
2022/07/01 12:43:00 client.go:112: -----------------------------------
2022/07/01 12:43:00 client.go:113: Disk IO :  100000  KB/s
2022/07/01 12:43:00 client.go:114: Net speed:  125000  KB/s
2022/07/01 12:43:00 client.go:115: Expect memory size:  36571.42857142857 KB
2022/07/01 12:43:00 client.go:116: -----------------------------------
2022/07/01 12:44:12 client.go:169: ----------------volume----------------------
2022/07/01 12:44:12 client.go:170: data-size(KB) :  116 	 transfer time(s):  0.556464643
2022/07/01 12:44:12 client.go:171: --------------------------------------------
2022/07/01 12:44:12 client.go:250: ---------------------dump------------------------
2022/07/01 12:44:12 client.go:251: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/07/01 12:44:12 client.go:252: 1.006370963 	 377820 	 8.082280512
2022/07/01 12:44:12 client.go:253: -------------------------------------------------
2022/07/01 12:44:14 client.go:266: The downtime is  11.64298908s
2022/07/01 12:44:14 client.go:274: The total migration time is  1m26.098747947s
2022/07/01 12:44:14 client.go:29: ---------------------PrintInfo--------------------------------------
2022/07/01 12:44:14 client.go:30: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/07/01 12:44:14 client.go:32: 0 	 182544 		 0.267510616 	 4.714376081
2022/07/01 12:44:14 client.go:32: 1 	 97928 		 0.359147977 	 3.518005315
2022/07/01 12:44:14 client.go:32: 2 	 122504 		 0.366676166 	 4.11595788
2022/07/01 12:44:14 client.go:32: 3 	 105600 		 0.464673823 	 3.937125285
2022/07/01 12:44:14 client.go:32: 4 	 192760 		 0.490771902 	 5.727287775
2022/07/01 12:44:14 client.go:32: 5 	 168744 		 0.532307117 	 6.026937439
2022/07/01 12:44:14 client.go:32: 6 	 239320 		 0.630361296 	 6.727922201
2022/07/01 12:44:14 client.go:32: 7 	 194560 		 0.755242443 	 6.440943812
2022/07/01 12:44:14 client.go:32: 8 	 301456 		 0.71302023 	 8.015583913
2022/07/01 12:44:14 client.go:32: 9 	 244472 		 0.911295092 	 8.047872389
2022/07/01 12:44:14 client.go:34: --------------------------------------------------------------------
```
```0
369M	./checkpoint
179M	./checkpoint0
96M ./checkpoint1
120M	./checkpoint2
104M	./checkpoint3
189M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```1
369M	./checkpoint
105M	./checkpoint0
96M	./checkpoint1
120M	./checkpoint2
104M	./checkpoint3
189M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```2
369M	./checkpoint
66M	./checkpoint0
35M	./checkpoint1
120M	./checkpoint2
104M	./checkpoint3
189M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```3
369M	./checkpoint
51M	./checkpoint0
25M	./checkpoint1
67M	./checkpoint2
104M	./checkpoint3
189M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```4
369M	./checkpoint
39M	./checkpoint0
17M	./checkpoint1
16M	./checkpoint2
23M	./checkpoint3
189M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```5
369M	./checkpoint
28M	./checkpoint0
13M	./checkpoint1
12M	./checkpoint2
18M	./checkpoint3
78M	./checkpoint4
165M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```6
369M	./checkpoint
21M	./checkpoint0
10M	./checkpoint1
9.4M	./checkpoint2
14M	./checkpoint3
22M	./checkpoint4
35M	./checkpoint5
234M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```7
369M	./checkpoint
16M	./checkpoint0
7.9M	./checkpoint1
7.4M	./checkpoint2
11M	./checkpoint3
18M	./checkpoint4
21M	./checkpoint5
110M	./checkpoint6
191M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```8
369M	./checkpoint
12M	./checkpoint0
6.4M	./checkpoint1
6.0M	./checkpoint2
8.6M	./checkpoint3
16M	./checkpoint4
15M	./checkpoint5
19M	./checkpoint6
35M	./checkpoint7
295M	./checkpoint8
239M	./checkpoint9
```

```9
369M	./checkpoint
9.2M	./checkpoint0
5.3M	./checkpoint1
5.0M	./checkpoint2
7.1M	./checkpoint3
14M	./checkpoint4
11M	./checkpoint5
13M	./checkpoint6
22M	./checkpoint7
129M	./checkpoint8
239M	./checkpoint9
```

```10
369M	./checkpoint
7.3M	./checkpoint0
4.7M	./checkpoint1
4.4M	./checkpoint2
6.3M	./checkpoint3
13M	./checkpoint4
8.3M	./checkpoint5
8.9M	./checkpoint6
16M	./checkpoint7
17M	./checkpoint8
36M	./checkpoint9
```

## idle_redis motivation
```
2022/07/01 14:31:30 server.go:52: idle_redis
2022/07/01 14:31:39 server.go:75: restore:idle_redis
2022/07/01 14:31:39 server.go:105: Restore time is  79.032776ms
2022/07/01 14:31:39 server.go:108: Handle finished.
```

```
2022/07/01 14:31:30 client.go:149: -----------------config.json------------------
2022/07/01 14:31:30 client.go:150: data-size(KB) :  12 	 transfer time(s):  0.569620308
2022/07/01 14:31:30 client.go:151: ----------------------------------------------
2022/07/01 14:31:37 client.go:157: --------------------rootfs--------------------
2022/07/01 14:31:37 client.go:158: data-size(KB) :  125280 	 transfer time(s):  6.615828038
2022/07/01 14:31:37 client.go:159: ----------------------------------------------
2022/07/01 14:31:37 client.go:112: -----------------------------------
2022/07/01 14:31:37 client.go:113: Disk IO :  100000  KB/s
2022/07/01 14:31:37 client.go:114: Net speed:  125000  KB/s
2022/07/01 14:31:37 client.go:115: Expect memory size:  36571.42857142857 KB
2022/07/01 14:31:37 client.go:116: -----------------------------------
2022/07/01 14:31:39 client.go:169: ----------------volume----------------------
2022/07/01 14:31:39 client.go:170: data-size(KB) :  116 	 transfer time(s):  0.445555121
2022/07/01 14:31:39 client.go:171: --------------------------------------------
2022/07/01 14:31:39 client.go:250: ---------------------dump------------------------
2022/07/01 14:31:39 client.go:251: dumpTime(s)	 data-size(KB)	 transfer time(s)
2022/07/01 14:31:39 client.go:252: 0.035646759 	 464 	 0.453022403
2022/07/01 14:31:39 client.go:253: -------------------------------------------------
2022/07/01 14:31:39 client.go:266: The downtime is  1.016689482s
2022/07/01 14:31:39 client.go:274: The total migration time is  8.863070147s
2022/07/01 14:31:39 client.go:29: ---------------------PrintInfo--------------------------------------
2022/07/01 14:31:39 client.go:30: index	 data-size(KB)		 pre-time(s)	 transfer-time(s)
2022/07/01 14:31:39 client.go:32: 0 	 2196 		 0.027699627 	 0.610223838
2022/07/01 14:31:39 client.go:34: --------------------------------------------------------------------
```