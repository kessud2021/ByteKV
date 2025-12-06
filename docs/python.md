
# Python Client Usage

```
from pybytekv import ByteKVClient

client = ByteKVClient("127.0.0.1", 6379)
```

# Set and Get
```
client.set("foo", "bar")
print(client.get("foo"))
```

# TTL / Expire
```
client.set("temp", "123", ttl=10)
print(client.ttl("temp"))
client.expire("temp", 30)
```

# Delete key
```
client.delete("foo")
```

# Ping
```
print(client.ping())
```

# Publish
```
client.publish("news", "Hello World")
```
For full API reference, see Commands
