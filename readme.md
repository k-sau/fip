# fip

Takes list of IPs from stdin and filters private IP using golang's IP.IsPrivate() & IsLoopback()

### Installation
```
GO111MODULE=on go get github.com/k-sau/fip
```

### Usage
```
cat ips.txt | fip
```

```
cat ips.txt | fip

64.6.64.6
94.140.14.140
208.67.222.222
8.8.8.8
195.46.39.39
64.6.65.6
1.0.0.1
```
