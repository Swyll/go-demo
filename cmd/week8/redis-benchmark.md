## Redis对不同字节大小的Value的测试结果

### 不同字节的get和set的

#### 10字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 10 -t get,set -q <br>
SET: 94055.68 requests per second, p50=0.319 msec                   <br>
GET: 92626.91 requests per second, p50=0.327 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 10 -t get,set -q <br>
SET: 95356.15 requests per second, p50=0.311 msec                   <br>
GET: 97494.39 requests per second, p50=0.295 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 10 -t get,set -q <br>
SET: 96852.30 requests per second, p50=0.295 msec                   <br>
GET: 92954.08 requests per second, p50=0.319 msec                   <br>

### 20字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 20 -t get,set -q <br>
SET: 94250.71 requests per second, p50=0.311 msec                   <br>
GET: 95904.86 requests per second, p50=0.311 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 20 -t get,set -q <br>
SET: 97370.98 requests per second, p50=0.295 msec                   <br>
GET: 97799.51 requests per second, p50=0.287 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 20 -t get,set -q <br>
SET: 96292.73 requests per second, p50=0.303 msec                   <br>
GET: 95721.27 requests per second, p50=0.303 msec                   <br>

### 50字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 50 -t get,set -q <br>
SET: 95447.16 requests per second, p50=0.303 msec                   <br>
GET: 96292.73 requests per second, p50=0.303 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 50 -t get,set -q <br>
SET: 89031.34 requests per second, p50=0.343 msec                   <br>
GET: 97238.42 requests per second, p50=0.295 msec                   <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 50 -t get,set -q <br>
SET: 97694.41 requests per second, p50=0.295 msec                   <br>
GET: 97962.38 requests per second, p50=0.295 msec                   <br>

### 100字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 100 -t get,set -q <br>
SET: 91149.40 requests per second, p50=0.327 msec                    <br>
GET: 90122.57 requests per second, p50=0.327 msec                    <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 100 -t get,set -q <br>
SET: 95383.44 requests per second, p50=0.311 msec                    <br>
GET: 95565.75 requests per second, p50=0.311 msec                    <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 100 -t get,set -q <br>
SET: 95201.83 requests per second, p50=0.303 msec                    <br>
GET: 91124.48 requests per second, p50=0.335 msec                    <br>

### 200字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 200 -t get,set -q <br>
SET: 93075.20 requests per second, p50=0.311 msec                    <br>
GET: 94428.70 requests per second, p50=0.311 msec                    <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 200 -t get,set -q <br>
SET: 88035.91 requests per second, p50=0.343 msec                    <br>
GET: 94616.33 requests per second, p50=0.311 msec                    <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 200 -t get,set -q <br>
SET: 89293.69 requests per second, p50=0.335 msec                    <br>
GET: 92729.97 requests per second, p50=0.311 msec                    <br>

### 1k字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 1000 -t get,set -q <br>
SET: 90628.96 requests per second, p50=0.327 msec                     <br>
GET: 90391.40 requests per second, p50=0.319 msec                     <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 1000 -t get,set -q <br>
SET: 91365.92 requests per second, p50=0.311 msec                     <br>
GET: 95465.39 requests per second, p50=0.311 msec                     <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 1000 -t get,set -q <br>
SET: 88519.08 requests per second, p50=0.335 msec                     <br>
GET: 89637.86 requests per second, p50=0.327 msec                     <br>

### 5k字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 5000 -t get,set -q <br>
SET: 85266.03 requests per second, p50=0.311 msec                     <br>
GET: 84904.06 requests per second, p50=0.327 msec                     <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 5000 -t get,set -q <br>
SET: 83243.16 requests per second, p50=0.327 msec                     <br>
GET: 81686.00 requests per second, p50=0.343 msec                     <br>

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 5000 -t get,set -q <br> 
SET: 82054.65 requests per second, p50=0.327 msec                     <br>
GET: 83070.27 requests per second, p50=0.335 msec                     <br>

### 50k字节

root@c991e4665b95:/# redis-benchmark -n 1000000 -d 50000 -t get,set -q <br>
SET: 39758.27 requests per second, p50=0.855 msec                      <br>
GET: 32208.19 requests per second, p50=0.735 msec                      <br>

### 总结，在字节小于1000字节的时候redis的get和set的qps波动不大，到了5000字节甚至之上qps波动直接腰斩

## Redis对不同字节大小的Value的key的大小

### 1W字节
[root@swy redis]# redis-memory-for-key -s 127.0.0.1 1
Key				1
Bytes				10280
Type				string

### 10W字节
[root@swy redis]# redis-memory-for-key -s 127.0.0.1 2
Key				2
Bytes				114728
Type				string

### 50W字节
[root@swy redis]# redis-memory-for-key -s 127.0.0.1 2
Key				2
Bytes				524328
Type				string


