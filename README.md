<div align="center">

# Statis

**A simple, clean status page and lightweight monitoring service.**

[![CodeFactor](https://www.codefactor.io/repository/github/jesperrichert/statis/badge)](https://www.codefactor.io/repository/github/jesperrichert/statis)

</div>

---

## What is Statis?

Third time's the charm: Statis is a simple and clean status page with a small monitoring service built in. It comes with
a tidy UI and a simple yet powerful API that lets you customize the whole application.

- Change the design
- Remove the whole frontend
- Build any application on top of the Statis API (Discord bot, mobile app, ...)

## Features

| Feature                     | Description                                       |
|-----------------------------|---------------------------------------------------|
| **Auto configuration**      | Configure everything via a JSON config (URL)      |
| **Webhooks**                | Share your monitoring data with any other service |
| **Clean UI**                | View all your monitors at a glance                |
| **Incidents & maintenance** | Announce upcoming maintenance windows             |
| **And more**                | Some other small features ;)                      |

## Installation

Use the provided Docker Compose file: [`compose.yml`](compose.yml)

Set up what you need and let Statis do the rest.

> **API docs:** The API documentation lives in the [`api-specs`](./docs) branch. <br />
> *Also you can find more about Statis at [@Outline](https://doc.jespersen.zip/s/statis)*

## Contributing

Feature requests are always welcome! Things I'd like to add:

- [ ] OpenTelemetry data support for display !!
- [ ] More monitor types !
- [ ] Badges !!!
- [ ] Cli :) !!
- [ ] *\<your idea here\>* 