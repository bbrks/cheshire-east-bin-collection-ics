# cheshire-east-bin-collection-ics

Collects Cheshire East bin collection schedule for a Unique Property Reference Number (UPRN) and serves an auto-updating `.ics` calendar at a public URL.

## Prerequisites

### Cheshire East UPRN

You can find your Unique Property Reference Number (UPRN) by following these steps:

1. Use the `Address Finder` on the [Cheshire East Public Map Viewer](https://maps.cheshireeast.gov.uk/ce/webmapping)

2. Search for your address, and note your UPRN displayed

## Installation

### Running with Docker

    UPRN="100012357047" docker run -d -p 8080:8080 bbrks/cheshire-east-bin-collection-ics

### Running standalone

1. Download the latest release binary
 
2. Run the server with the UPRN:

       ./cheshire-east-bin-collection-ics -uprn="100012357047"

3. Subscribe to the calendar at the following URL:

       http://yourhost:8080/collections.ics

### Configuration Options

* `-uprn="100012357047"`
  * Set the UPRN
* `-update=24`
  * Set the update interval (in hours)
* `-addr="0.0.0.0:1234"`
  * Change the listen address/port