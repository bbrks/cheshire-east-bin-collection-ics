# 🗑📅 Cheshire East Bin Collection iCalendar (.ics)

[![Build Status](https://travis-ci.com/bbrks/cheshire-east-bin-collection-ics.svg)](https://travis-ci.com/bbrks/cheshire-east-bin-collection-ics) [![GitHub tag](https://img.shields.io/github/tag/bbrks/cheshire-east-bin-collection-ics.svg)](https://github.com/bbrks/cheshire-east-bin-collection-ics/releases) [![](https://img.shields.io/docker/pulls/bbrks/cheshire-east-bin-collection-ics.svg)](https://hub.docker.com/u/bbrks/cheshire-east-bin-collection-ics) [![license](https://img.shields.io/github/license/bbrks/cheshire-east-bin-collection-ics.svg)](https://github.com/bbrks/cheshire-east-bin-collection-ics/blob/master/LICENSE)

Fetches Cheshire East bin collection schedule for a Unique Property Reference Number (UPRN) and serves an auto-updating `.ics` calendar at a public URL.

**Disclaimer**: This project is pulling data directly from Cheshire East from a non-public API, and may be unstable, or even disappear without notice.

To prevent abuse, there are a few built-in mechanisms to prevent needlessly hammering the upstream data source by default (do bin collections update _that_ frequently?!)

## TOC

- [Prerequisites](#prerequisites)
  - [Cheshire East UPRN Lookup](#cheshire-east-uprn-lookup)
- [Installation](#installation)
  - [Running with Docker](#running-with-docker)
  - [Running standalone](#running-standalone)
- [Configuration Options](#configuration-options)

---

## Prerequisites

### Cheshire East UPRN Lookup

You can find your Unique Property Reference Number (UPRN) in Cheshire East by following these steps:

1. Use the `Address Finder` on the [Cheshire East Public Map Viewer](https://maps.cheshireeast.gov.uk/ce/webmapping)

![Cheshire East Public Map Viewer - Address Finder](https://i.imgur.com/9Ds72nV.png)

2. Search for your address, and note your UPRN

![Cheshire East Unique Property Reference Number UPRN](https://i.imgur.com/4sTjCpZ.png)

## Installation

### Running with Docker

    docker run -d -e UPRN="100012357047" -p 8080:8080 bbrks/cheshire-east-bin-collection-ics

### Running standalone

1. Download the [latest release binary](https://github.com/bbrks/cheshire-east-bin-collection-ics/releases)
 
2. Run the service with your UPRN:

       ./cheshire-east-bin-collection-ics -uprn="100012357047"

3. Your public calendar URL will be logged on startup:

        2019-05-28T23:32:32.614+01:00 [ALL] Serving calendar at http://[::]:8080/collections.ics
        2019-05-28T23:32:34.334+01:00 [INF] req:1 <-- GET /collections.ics from [::1]:58626
        2019-05-28T23:32:34.334+01:00 [INF] req:1 Fetching a fresh copy of data
        2019-05-28T23:32:39.325+01:00 [INF] req:1 --> 200 (OK) in 4.991430799s
        2019-05-28T23:32:42.235+01:00 [INF] req:2 <-- GET /collections.ics from [::1]:58660
        2019-05-28T23:32:42.235+01:00 [DBG] req:2 Skipping update and serving data from cache
        2019-05-28T23:32:42.237+01:00 [INF] req:2 --> 200 (OK) in 1.65305ms

## Configuration Options

- `-uprn="100012357047"`
  - Set the UPRN to fetch the collection schedule for
  - This is intentionally not parameterised in the public URL, to prevent accidentally hosting a public service 🙃
- `-updateInterval=24h`
  - Set how often the service looks for new data
  - Range: `24h` to `7d`
- `-addr="0.0.0.0:1234"`
  - Change the listen address/port

## Building from source

### Dockerfile

There's a multi-stage dockerfile in the repo which will compile and produce a lightweight runnable image:

    docker build -t cheshire-east-bin-collection-ics:latest .

### Binaries

With Go already installed, you can build a binary with the following command:

    go build .

## Contributing

Issues, feature requests or improvements welcome!

## Licence

This project is licensed under the [MIT License](LICENSE).
