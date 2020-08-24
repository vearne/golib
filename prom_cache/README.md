### prom_cache
对github.com/gogf/gf 提供的local cache进行了简单的封装
以实现自动添加local cache的prometheus指标
#### 用法
```
FixedCache = prom_cache.NewCacheWithProm("channel", 10000)
```
#### prometheus指标
```
# HELP cache_request_total Total number of local cache request
# TYPE cache_request_total counter
cache_request_total{kind="channel",state="all"} 29
cache_request_total{kind="channel",state="hit"} 20
# HELP cache_size size of local cache
# TYPE cache_size gauge
cache_size{kind="channel"} 1
```


