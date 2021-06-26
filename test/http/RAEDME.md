# HTTP APIs Test

## Functional Testing

We choose [curl](https://github.com/curl/curl) which is a CLI tool in Linux for transferring data with URL syntax as HTTP client to test out HTTP APIs.


### Set key-value pair

Send Request：

```Shell
curl -v 127.0.0.1:10615/cache/1 -XPUT -dyobol
```

Response:

```
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 10615 (#0)
> PUT /cache/1 HTTP/1.1
> Host: 127.0.0.1:10615
> User-Agent: curl/7.58.0
> Accept: */*
> Content-Length: 5
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 5 out of 5 bytes   
< HTTP/1.1 200 OK
< Date: Sat, 26 Jun 2021 13:55:33 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact 
```

### Get value by key

Send Request:

```shell
curl 127.0.0.1:10615/cache/1
```

Response:

```
yobol
```

### Delete value by key

Send Request：

```shell
curl 127.0.0.1:10615/cache/1 -XDELETE
```

Response:

```
```

### Get status

Send Request:

```shell
curl 127.0.0.1:10615/status
```

Response:

```
{"pairs":0,"key_size":0,"value_size":0}
```

## Performance Testing

### Redis benchmark

```shell
# install Redis in Debain OS
sudo apt-get install redis-server

# start Redis server
sudo /etc/init.d/redis-server start
# Starting redis-server: redis-server.

# execute the following commands in another terminal
redis-benchmark --help
# Usage: redis-benchmark [-h <host>] [-p <port>] [-c <clients>] [-n <requests>] [-k <boolean>]

#  -h <hostname>      Server hostname (default 127.0.0.1)
#  -p <port>          Server port (default 6379)
#  -s <socket>        Server socket (overrides host and port)
#  -a <password>      Password for Redis Auth
#  -c <clients>       Number of parallel connections (default 50)
#  -n <requests>      Total number of requests (default 100000)
#  -d <size>          Data size of SET/GET value in bytes (default 3)
#  --dbnum <db>       SELECT the specified db number (default 0)
#  -k <boolean>       1=keep alive 0=reconnect (default 1)
#  -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD
#   Using this option the benchmark will expand the string __rand_int__
#   inside an argument with a 12 digits number in the specified range
#   from 0 to keyspacelen-1. The substitution changes every time a command
#   is executed. Default tests use this to hit random keys in the
#   specified range.
#  -P <numreq>        Pipeline <numreq> requests. Default 1 (no pipeline).
#  -e                 If server replies with errors, show them on stdout.
#                     (no more than 1 error per second is displayed)
#  -q                 Quiet. Just show query/sec values
#  --csv              Output in CSV format
#  -l                 Loop. Run the tests forever
#  -t <tests>         Only run the comma separated list of tests. The test
#                     names are the same as the ones produced as output.
#  -I                 Idle mode. Just open N idle connections and wait.

# Examples:

#  Run the benchmark with the default configuration against 127.0.0.1:6379:
#    $ redis-benchmark

#  Use 20 parallel clients, for a total of 100k requests, against 192.168.1.1:
#    $ redis-benchmark -h 192.168.1.1 -p 6379 -n 100000 -c 20

#  Fill 127.0.0.1:6379 with about 1 million keys only using the SET test:
#    $ redis-benchmark -t set -n 1000000 -r 100000000

#  Benchmark 127.0.0.1:6379 for a few commands producing CSV output:
#    $ redis-benchmark -t ping,set,get -n 100000 --csv

#  Benchmark a specific command line:
#    $ redis-benchmark -r 10000 -n 10000 eval 'return redis.call("ping")' 0

#  Fill a list with 10000 random elements:
#    $ redis-benchmark -r 10000 -n 10000 lpush mylist __rand_int__

#  On user specified command lines __rand_int__ is replaced with a random integer       
#  with a range of values selected by the -r option.

redis-benchmark -c 1 -t set,get -n 100000 -d 1000 -r 100000
# ====== SET ======
#   100000 requests completed in 7.38 seconds
#   1 parallel clients
#   1000 bytes payload
#   keep alive: 1

# 100.00% <= 1 milliseconds
# 100.00% <= 3 milliseconds
# 13550.14 requests per second

# ====== GET ======
#   100000 requests completed in 6.93 seconds
#   1 parallel clients
#   1000 bytes payload
#   keep alive: 1

# 100.00% <= 0 milliseconds
# 14430.01 requests per second

# stop Redis server
sudo /etc/init.d/redis-server stop
# Stopping redis-server: redis-server
```

### go-cache benchmark

```shell
# install go-redis
go get -u github.com/go-redis/redis

cd ../benchmark
go build .   # generate benchmark.exe in Windows or benchmark in Linux
.\benchmark.exe --help
# Usage of D:\Project\Go\go-cache\test\benchmark\benchmark.exe:
#   -P int
#         pipeline length (default 1)
#   -c int
#         number of parallel connections (default 1)
#   -d int
#         data size of SET/GET value in bytes (default 1000)   
#   -h string
#         cache server address (default "localhost")
#   -n int
#         total number of requests (default 1000)
#   -r int
#         keyspacelen, use random keys from 0 to keyspacelen-1
#   -t string
#         test set, could be get/set/mixed (default "set")
#   -type string
#         cache server type (default "redis")

.\benchmark.exe -type http -n 100000 -r 100000 -t mixed # Windows
# type is http
# server is localhost
# total 100000 requests
# data size is 1000
# we have 1 connections
# operation is set
# keyspacelen is 100000
# pipeline length is 1
# 0 records get
# 0 records miss
# 100000 records set
# 7.248507 seconds total
# 99% requests < 1 ms
# 99% requests < 2 ms
# 100% requests < 8 ms
# 70 usec average for each request
# throughput is 13.795944 MB/s
# rps is 13795.944271
.\benchmark.exe -type http -n 100000 -r 100000 -t get # Windows
# type is http
# server is localhost  
# total 100000 requests
# data size is 1000    
# we have 1 connections
# operation is get     
# keyspacelen is 100000
# pipeline length is 1 
# 0 records get
# 100000 records miss
# 0 records set
# 7.157247 seconds total
# 99% requests < 1 ms
# 99% requests < 2 ms
# 100% requests < 7 ms
# 69 usec average for each request
# throughput is 0.000000 MB/s
# rps is 13971.852000
```

Colusion: If we use HTTP as the transferring protocol, the processing speed is about one-tenth of Redis.