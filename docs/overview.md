``markdown
# Project Overview

ByteKV consists of three main components:

1. **ByteKV Server** – In-memory TCP server handling key-value storage, TTL, PUBLISH, and RESP commands.
2. **Python Client (`pybytekv`)** – Connects via TCP; supports standard commands.
3. **Node.js Client (`bytekv`)** – Async Promise-based client with the same commands.

---

### ASCII Diagram

  ┌─────────────┐
  │ ByteKV      │
  │ TCP Server  │
  └─────┬───────┘
        │ RESP
        │
┌──────────┴───────────┐
│                      │
│                      │
Python Client  Node.js Client
(pybytekv) (bytekv)
