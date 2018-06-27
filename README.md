# Sensu Graphite Handler

The Sensu Graphite Handler is a [Sensu Event Handler][3] that sends metrics to
the time series database [Graphite][2]. [Sensu][1] can collect metrics using
check output metric extraction or the StatsD listener. Those collected metrics
pass through the event pipeline, allowing Sensu to deliver the metrics to the
configured metric event handlers.

## Installation

Download the latest version of the sensu-graphite-handler from [releases][4],
or create an executable script from this source.

From the local path of the sensu-influxdb-handler repository:
```
go build -o /usr/local/bin/sensu-graphite-handler main.go
```

## Configuration

Example Sensu 2.x handler definition:
```
{
  "name": "graphite",
  "type": "pipe",
  "command": "sensu-graphite-handler --host '123.4.5.6' --port 2003 --verbose"
}
```

Example Sensu 2.x check definition:
```
{
  "name": "collect-metrics",
  "command": "collect.sh",
  "interval": 10,
  "subscriptions": [
    "system"
  ],
  "output_metric_format": "graphite_plaintext",
  "output_metric_handlers": [graphite]
}
```

## Usage Examples

Help:
```
Usage:
  sensu-graphite-handler [flags]

Flags:
  -h, --help          help for sensu-graphite-handler
      --host string   the graphite carbon host (default "127.0.0.1")
  -p, --port int      the graphite carbon tcp port (default 2003)
  -t, --timeout int   the graphite carbon write timeout (default 5)
  -v, --verbose       verbose handler output

[1]: https://github.com/sensu/sensu-go
[2]: https://graphiteapp.org/
[3]: https://docs.sensu.io/sensu-core/2.0/reference/handlers/#how-do-sensu-handlers-work
[4]: https://github.com/portertech/sensu-graphite-handler/releases
