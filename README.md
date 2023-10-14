# Check Praivate IP

```bash
hostname -I
```

# Request Key Value

```bash
curl "http://20.0.0.10/?key=${YOUR_KEY}&value=${YOUR_VALUE}"
```

# Redis

```bash
redis-cli --raw -h 20.0.0.20
20.0.0.20:6379> keys *
20.0.0.20:6379> get test
```
