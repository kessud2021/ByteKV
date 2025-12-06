# Node.js Client Usage

```javascript
import { ByteKVClient } from "bytekv";

const client = new ByteKVClient("127.0.0.1", 6379);

// Set and Get
await client.set("foo", "bar");
console.log(await client.get("foo"));

// TTL / Expire
await client.expire("temp", 30);
console.log(await client.ttl("temp"));

// Delete key
await client.del("foo");

// Ping
console.log(await client.ping());

// Publish
await client.publish("news", "Hello World");
```

For full API reference, see Commands
