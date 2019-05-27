# cheshire-east-bin-collection-ics

Fetches Cheshire East bin collection schedule for a Unique Property Reference Number (UPRN) and serves an auto-updating `.ics` calendar at a public URL.

- [Prerequisites](#prerequisites)
  - [Cheshire East UPRN Lookup](#cheshire-east-uprn-lookup)
- [Installation](#installation)
  - [Running with Docker](#running-with-docker)
  - [Running standalone](#running-standalone)
- [Configuration Options](#configuration-options)

## Prerequisites

### Cheshire East UPRN Lookup

You can find your Unique Property Reference Number (UPRN) in Cheshire East by following these steps:

1. Use the `Address Finder` on the [Cheshire East Public Map Viewer](https://maps.cheshireeast.gov.uk/ce/webmapping)

![Cheshire East Public Map Viewer - Address Finder](https://i.imgur.com/9Ds72nV.png)

2. Search for your address, and note your UPRN displayed

![Cheshire East Unique Property Reference Number](https://i.imgur.com/4sTjCpZ.png)

## Installation

### Running with Docker

    UPRN="100012357047" docker run -d -p 8080:8080 bbrks/cheshire-east-bin-collection-ics

### Running standalone

1. Download the latest release binary
 
2. Run the server with the UPRN:

       ./cheshire-east-bin-collection-ics -uprn="100012357047"

3. Subscribe to the calendar at the following URL:

       http://yourhost:8080/collections.ics

## Configuration Options

* `-uprn="100012357047"`
  * Set the UPRN
* `-update=24`
  * Set the update interval (in hours)
* `-addr="0.0.0.0:1234"`
  * Change the listen address/port
