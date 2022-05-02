# catalystsquad-platform Pulumi Component Provider (Go)

TODO explain things

## Prerequisites

- Go 1.15
- Pulumi CLI
- Node.js (to build the Node.js SDK)
- Yarn (to build the Node.js SDK)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)

## Build and Test

```bash
# Build and install the provider (plugin copied to $GOPATH/bin)
make install_provider

# Regenerate SDKs
make generate

# Test Go SDK
$ make install_go_sdk
$ cd examples/simple-vpc-go
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up
```

# TODO

- [ ] document components
- [ ] example implementation
