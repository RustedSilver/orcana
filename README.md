
# 🚧 WIP: orcana

> **Note:** This is a personal project, currently in its early stages.  
> I'm building it by myself in an unfamliar languange — expect rough edges, possible bugs, and changes without warning.
> I at least guaranty that any commit of mine will be a tested functionality

## 🧠 What is this?

**Orcana** is an experimental tool I'm working on **to add additional functionality on a docker compose orchestration**.  
It's mainly intended for use in **single-node environment like, a single server on top of docker compose**, and it may or may not evolve into something more robust.

## 🔍 Why?

Because I repurposed an old PC into a home server and started running a local infrastructure using Docker, I wanted a simpler way to manage my Docker Compose setup. Kubernetes and Swarm felt like overkill for a single-node environment, so I set out to build something lightweight that fit my needs.
I couldn’t find an existing tool that did exactly what I had in mind, so I decided to roll my own — and it's also a great excuse to dive deeper into Go and learn more about building CLI applications.

## 🚀 Current Status

Very immature. It kind of works, sometimes. I'm still figuring out:

- Requires **Docker** and **Docker Compose** (the newer `docker compose` CLI).  
  It **won't work** with the legacy `docker-compose` binary.
- Currently, the only feature is the ability to **enable or disable services** within a Compose file — similar in spirit to how Helm manages components in Kubernetes.
- The goal is for this to evolve into a more **mature tool for managing local servers**.
  At the moment, it simply wraps `docker compose` using `exec.Command`, but I'm exploring better ways to integrate with the Docker API in the future.

## 🛠 Tech Stack

- go/cobra
- go-yaml
- Whatever else I feel like adding as I go

## 📦 Installation (if it even makes sense right now)

```bash
# warning: this might not work yet
git clone https://github.com/RustedSilver/orcana.git
cd orcana
make build

cp ./bin/orcana /usr/local/bin/orcana
```


