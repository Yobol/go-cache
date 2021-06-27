# go-cache

A experimental cache system is built to study.

## APIs

### HTTP

### TCP

We choose ABNF definited in[RFC 5234](https://datatracker.ietf.org/doc/html/rfc5234) to describe our APIs for TCP.

```
command = op key | key-value
op = 'S' | 'G' | 'D'
key = bytes-array
bytes-array = length SP content
length = 1*DIGIT
content = *OCTET
key-value = length SP length SP content content
response = error | bytes-array
error = '-' bytes-array
```

`PS1: SP meaning SPACE, DIGIT meaning 0~9 and OCTET meaning 0x00~0xFF are all defined in ABNF.`
`PS2: Space ' ' in rule is not actually existed.`

#### SET

Request:

```
S<klen><SP><vlen><SP><key><value>
```

Successful Response:

```
0
```

Error Response:

```
-<error>
```

#### GET

Request:

```
G<klen><SP><key>
```

Successful Response:

```
<vlen><SP><value>
```

Error Response:

```
-<error>
```

#### DEL

Request:

```
D<klen><SP><key>
```

Successful Response:

```
0
```

Error Response:

```
-<error>
```