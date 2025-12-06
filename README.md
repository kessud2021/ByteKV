
# **ByteKV**

![PyPI](https://img.shields.io/pypi/v/pybytekv?color=blue\&label=Python)
![npm](https://img.shields.io/npm/v/bytekv?color=green\&label=Node.js)
![Coverage](https://img.shields.io/codecov/c/github/VSS-CO/bytekv?label=Coverage)
![License](https://img.shields.io/badge/license-MIT-lightgrey)

ByteKV is a fast, TCP-based in-memory key-value store with a Redis-like RESP protocol. It has official clients for **Python** and **Node.js**, allowing you to interact with the server from either ecosystem.

---

## **Installation**

* **Python**

```bash
pip install pybytekv
```

* **Node.js**

```bash
npm install bytekv
```

---

## **Quick Start**

### Python Usage

```python
from pybytekv import ByteKVClient

client = ByteKVClient("127.0.0.1", 6379)

# Set a key
client.set("foo", "bar")

# Get a key
print(client.get("foo"))

# TTL commands
client.set("temp", "123", ttl=10)
print(client.ttl("temp"))
client.expire("temp", 30)

# Delete
client.delete("foo")

# Ping
print(client.ping())

# Publish
client.publish("news", "Hello World")
```

### Node.js Usage

```javascript
import { ByteKVClient } from "bytekv";

const client = new ByteKVClient("127.0.0.1", 6379);

await client.set("foo", "bar");
console.log(await client.get("foo"));

await client.expire("temp", 30);
console.log(await client.ttl("temp"));

await client.del("foo");
console.log(await client.ping());

await client.publish("news", "Hello World");
```

---

## **Project Overview**

ByteKV consists of three main components:

1. **ByteKV Server** – Redis-like in-memory TCP server handling key-value storage, TTL, PUBLISH, and RESP command parsing.
2. **Python Client (`pybytekv`)** – Connects via TCP to ByteKV server; supports all standard commands.
3. **Node.js Client (`bytekv`)** – Async, Promise-based client with the same commands as Python client.

---

### **ASCII Diagram**

```
      ┌─────────────┐
      │ ByteKV      │
      │ TCP Server  │
      └─────┬───────┘
            │ RESP
            │
 ┌──────────┴───────────┐
 │                      │
 │                      │
Python Client         Node.js Client
(pybytekv)             (bytekv)
```

---

### **Data Flow Example**

1. Python client sets a key → server stores it.
2. Node.js client reads the key → server returns the value.
3. TTL/EXPIRE commands update expiration times.
4. PUBLISH messages are sent to all connected subscribers.

---

## **Tips & Best Practices**

* Python and Node.js clients behave similarly — you can mix them in the same system.
* Always handle exceptions in Python (`try/except`) or Node.js (`try/catch`).
* Use TTL for temporary keys to avoid memory buildup.
* PUBLISH only delivers messages to currently connected subscribers.
* Keep clients updated to leverage all commands (`SET`, `GET`, `DEL`, `EXPIRE`, `TTL`, `PING`, `PUBLISH`).

---

## **Features**

* SET, GET, DEL operations
* TTL and EXPIRE for key expiration
* PING health check
* PUBLISH for basic messaging
* Works with Python 3.7+ and Node.js 14+

---

## **Contributing**

We welcome contributions! You can help by reporting issues, adding features, or improving documentation.

### **Python Client (`pybytekv`)**

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/bytekv.git
cd bytekv/pybytekv
```

2. **Create a virtual environment**

```bash
python -m venv venv
source venv/bin/activate        # macOS/Linux
venv\Scripts\activate           # Windows
```

3. **Install dependencies**

```bash
pip install -r requirements.txt
```

4. **Run tests**

```bash
python -m unittest discover tests
```

5. **Make changes**, then commit and push.

6. **Build and publish locally for testing**

```bash
python -m build
python -m pip install dist/pybytekv-<version>.tar.gz
```

---

### **Node.js Client (`bytekv`)**

1. **Navigate to Node.js client folder**

```bash
cd bytekv/node-client
```

2. **Install dependencies**

```bash
npm install
```

3. **Run tests**

```bash
npm test
```

4. **Build (if using TypeScript or bundler)**

```bash
npm run build
```

5. **Link locally for testing**

```bash
npm link
# In a test project
npm link bytekv
```

6. **Make changes**, commit, and push.

---

### **Tips for Contributors**

* Follow **semantic versioning** when adding features or fixing bugs.
* Update **README.md** if you add new commands or examples.
* For new features, write **unit tests** for both Python and Node.js clients.
* Ensure **cross-language compatibility** — commands should behave the same in both clients.

---

## **License**

MIT — free to use, modify, and distribute.

---
