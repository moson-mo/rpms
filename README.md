# rpms
## Renoir power metrics server
</br>

A service that creates an HTTP/REST endpoint allowing to query power / performance metrics for AMD Renoir processors.
</br>
Note that the following kernel module needs to be installed [ryzen_smu](https://gitlab.com/leogx9r/ryzen_smu/)
Installation instructions can be found there...
</br>

## Purpose

Purpose is to make power/performance metrics available in an easy manner.</br>
Metrics can be queried via simple HTTP calls.</br>
You can ,for example, use it in the famous conky system monitoring tool by doing some curl calls...</br>
Or create your own web application with some graphs.</br></br>

I will provide some examples and most probably also a ui application with graphs, etc.</br>

## How to install

#### Manual

Binaries are available from the [releases](https://github.com/moson-mo/rpms/releases) page.</br>
If you have go installed on your machine, you can install with `go install github.com/moson-mo/rpms`</br>
Once installed you can create a systemd service to run system-startup (see [rpms.service](https://github.com/moson-mo/rpms/blob/main/assets/rpms.service)).</br>

#### AUR package

If you're using Arch or an Arch-based Distribution, there is an [AUR package](https://aur.archlinux.org/packages/rpms/) available.</br>
Use your favorite AUR helper to install.</br>
F.e.: `yay -S rpms`
</br>

## How to build

* Install go from your package manager or download it from the [Golang](https://golang.org/dl/) site.
* Download with `go get github.com/moson-mo/rpms`
* Change to package dir: `cd $(go env GOPATH)/src/github.com/moson-mo/rpms/`
* Build with `go build`
</br>

## How to use

The program needs to be run with root permissions. I recommend running it as a systemd service.

#### HTTP endpoints

Endpoint | Method | Description
--- | --- | ---
/pmtab|GET|Returns full pm table in json format.</br>Use URL parameter `?format=plain` to get a plain text version.
/pmval?metric=xyz|GET|Returns plain text value for a certain metric.</br>Example: `/pmval?metric=SOCKET POWER`</br>Use `/pmtab` to get a full list of available metrics.

#### Arguments

Argument | Type | Description
--- | --- | ---
-acao|string|Sets the Access-Control-Allow-Origin header if you want to allow querying the API from a webserver.</br>The default value is `null` to allow queries from local resources like an html file. (default `null`)
-address|string|The network address for the HTTP server.</br>Define `any` to listen on all interfaces. (default `127.0.0.1`)
-interval|duration|Query interval for reading data from the PM table. (default `1s`)
-port|int|Port number for the REST API server. (default `8090`)
</br>

## Dependencies / Prerequisites

* [ryzen_smu](https://gitlab.com/leogx9r/ryzen_smu/) - Linux kernel driver that exposes access to the SMU (System Management Unit) for AMD processors
</br>

## Currently supported processors

* Ryzen 7 (8-core) Renoir APU's - 4xx0 U/G/H/HS series

<b>Please contribute and help to support more models:</b>
Ryzen 3 (4 core) and Ryzen 5 (6-core) Renoir would really be interesting.



Simply create a data dump of the pm table: <b>[Please post here](https://gitlab.com/leogx9r/ryzen_smu/-/issues/1)</b>

</br>

## Available metrics

The PM table has a huge amount of different metrics available.</br>
For example:

* CPU speed / frequency (per core)
* CPU power consumption (total / per core)
* CPU voltage (per core)
* CPU temperature (core, socket, gfx)
* CPU temperature limits (core, socket, gfx)
* CPU power limits (STAPM, SLOW, FAST...)
* CPU usage (total / per core)
* GPU speed / frequency
* GPU power consumption
* GPU voltage
* GPU temperature
* GPU usage</br>
and tons of other things... check [tables.go](https://github.com/moson-mo/rpms/blob/main/tables.go) for a list
</br>

## Thanks to

* [Leonardo Gates](https://gitlab.com/leogx9r/) for the work on the Ryzen SMU driver module
* [sbski](https://github.com/sbski) for his reverse engineering work on the PM tables
* The people contributing with PM table data dumps
* The ones I forgot to mention here :)
</br>