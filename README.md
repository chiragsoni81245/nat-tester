# net-tester

A lightweight CLI toolkit for exploring real-world network behavior, starting with **NAT type detection using STUN**.

Built from scratch in Go to understand how peer-to-peer systems (like WebRTC, NetBird, libp2p) actually work under the hood.

---

## Features

* 🔍 Detect NAT type (symmetric vs non-symmetric)
* 🌐 Query multiple STUN servers
* 🧪 Uses a single UDP socket (realistic NAT behavior)
* ⚙️ Custom STUN server support
* 🧠 Clean, extensible architecture for future P2P experiments

---

## Installation

1. Use Built binary in releases

https://github.com/chiragsoni81245/nat-tester/releases

2. Build your own binary

```bash
git clone https://github.com/chiragsoni81245/net-tester.git
cd net-tester

go mod tidy
make build
./bin/net-tester
```

### Custom STUN servers

```bash
./nat-tester \
  --stun=stun.l.google.com:19302,stun.cloudflare.com:3478 \
  --timeout=5 \
  -v
```

---

## Example Output

```text
Local address: [::]:54321

[+] Running STUN tests...
[+] stun.l.google.com:19302 → 49.x.x.x:62000
[+] stun1.l.google.com:19302 → 49.x.x.x:62000

========== RESULT ==========
NAT Type: Restricted Cone NAT (likely)
```

---

## NAT Types Explained

| NAT Type        | Meaning                      | P2P Feasibility |
| --------------- | ---------------------------- | --------------- |
| Open Internet   | No NAT                       | 🟢 Easy         |
| Full Cone       | No filtering                 | 🟢 Easy         |
| Restricted Cone | Allows known IPs only        | 🟢 Works        |
| Port Restricted | IP + port restricted         | 🟡 Medium       |
| Symmetric NAT   | Port changes per destination | 🔴 Hard         |

---

## How It Works

1. Opens a **single UDP socket**
2. Sends STUN binding requests to multiple servers
3. Compares returned public IP:port mappings
4. Infers NAT behavior:

   * Same port → non-symmetric NAT
   * Different ports → symmetric NAT
5. Outputs best-effort classification

---

## Limitations

* NAT detection is **heuristic-based**, not exact
* Public STUN servers (like Google) don’t fully support RFC 5780
* Real systems (WebRTC, NetBird) rely on **connection attempts**, not just classification

---

## Project Structure

```
net-tester/
├── cmd/
│   └── nat-tester/        # CLI entrypoint
├── internal/
│   ├── stunclient/        # STUN request logic
│   ├── detector/          # NAT detection logic
│   └── types/             # shared types
```

---

## Why This Exists

This project is part of a deeper goal:

> Build a **custom P2P transport layer** from scratch.

---
