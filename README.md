### pressure test toy for architecture training camp

#### build

default build for mac

```bash
cd pressure-test-toy
make
```

if you want build for linux:

```bash
cd pressure-test-toy
make linux
```

#### usage

```bash
./pttoy-darwin.v0.0.1.bin -url https://www.baidu.com -concurrentNum 10 -totalReqNum 100
```

output:

```log
avg response time:	0.08s
95% response time:	0.07s
```
