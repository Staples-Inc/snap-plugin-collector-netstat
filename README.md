# snap collector plugin - netstat
This plugin collects metrics from /proc/net to collect the number of TCP connections active on the host.

## Getting Started
### System Requirements
* [golang 1.5+](https://golang.org/dl/) - needed only for building

### Operating systems
All OSs currently supported by plugin:
* Linux/amd64

### To build the plugin binary:
Fork https://github.com/staples-inc/snap-plugin-collector-netstat
Clone repo into `$GOPATH/src/github.com/staples-inc/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-netstat.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

## Documentation
### Collected Metrics
List of collected metrics in [METRICS.md](https://github.com/staples-inc/snap-plugin-collector-netstat/blob/master/METRICS.md)

## Roadmap
* Allow for custom specification of /proc/ path

Please contact us with any suggestions to improve the plugin. Open an issue or pull request and we will be sure to get back to you.
