# TCP APIs Test

## Functional Testing

```shell
cd test\tcp

go build .

.\tcp.exe --help
# Usage of D:\Project\Go\go-cache\test\tcp\tcp.exe:
#   -c string
#         command, could be get/set/del (default "get")
#   -h string
#         cache server host (default "localhost")
#   -k string
#         key
#   -p string
#         cache server port (default "10616")
#   -v string
#         value
```

### Set key-value pair

Send Request：

```Shell
.\tcp.exe -c set -k 1 -v yobol
```

Response:

```
yobol
```

### Get value by key

Send Request:

```shell
.\tcp.exe -c get -k 1
```

Response:

```
yobol
```

### Delete value by key

Send Request：

```shell
.\tcp.exe -c del -k 1
```

Response:

```

```