
# 🚧 WIP: [orcana]

> **Note:** This is a personal project, currently in its early stages.  
> I'm building it by myself in an unfamliar languange — expect rough edges, possible bugs, and changes without warning.
> I at least guaranty that any commit of mine will be a tested functionality

## 🧠 What is this?

**[Orcana]** is an experimental tool I'm working on **[to add additional functionality on a docker compose orchestration]**.  
It's mainly intended for use in **[single-node environment like, a single server on top of docker compose]**, and it may or may not evolve into something more robust.

## 🔍 Why?

Because I modified an old PC into a home server and am running a local infrastructure with docker on it. I needed something that simplified **[my docker compose orchestration as I did not want to use kubernetes or swarm as they are overkil]**.  
I couldn’t find an existing tool that fit exactly what I wanted, so I decided to roll my own. It’s also an excuse to learn more about **[golang and building cli applications]**.

## 🚀 Current Status

Very immature. It kind of works, sometimes. I'm still figuring out:

- [It requires docker & docker compose preinstalled. Will not work with old version of compose "docker-compose"]
- [Till now the only thing it does is giving the possibility to enable dissable a services on a compose file, similiar to helm on kubernetes]
- [I want it to become something more mature for a local server. At the moment it runs docker compose as an exec Command. I might view other alternatives]

## 🛠 Tech Stack

- [go/cobra]
- [go-yaml]
- Whatever else I feel like adding as I go

## 📦 Installation (if it even makes sense right now)

```bash
# warning: this might not work yet
git clone https://github.com/RustedSilver/orcana.git
cd orcana
make build

cp ./bin/orcana /usr/local/bin/orcana
```


