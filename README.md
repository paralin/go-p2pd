go-libp2pd
==================

[![](https://img.shields.io/badge/project-libp2p-blue.svg?style=flat-square)](http://libp2p.io/)

> libp2p daemon with a control cli.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

```sh
$ go get -u -v github.com/paralin/go-p2pd/cmd/p2pd
```

Optionally:

```sh
$ go get -u -v github.com/paralin/go-p2pd/cmd/p2pdctl
```

## Usage

You can use the control commands prefixed with either `p2pd ctl` or `p2pdctl` depending on how you built the client.

### Root `p2pd` Usage

```
NAME:
   p2pd - p2pd daemon and cli.

USAGE:
   p2pd [global options] command [command options] [arguments...]

AUTHOR:
   Christian Stewart <christian@paral.in>

COMMANDS:
     ctl      Ctl contains all control commands.
     daemon   starts the p2pd daemon
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### `p2pdctl node` Usage

```
NAME:
   p2pdctl node - Contains all node-related subcommands.

USAGE:
   p2pdctl node command [command options] [arguments...]

COMMANDS:
     add     Adds a new node to the p2pd instance.
     start   Starts a previously created node.
     listen  Listen commands a started node to listen on an additional address.
     status  Status checks node's status.

OPTIONS:
   --help, -h  show help
```

## Getting Started

Here is an example:

```sh
# Start the p2pd daemon
$ ./p2pd daemon --data-path=./data &
# Create a node
$ ./p2pd node add test
# Start it
$ ./p2pd node start test
# Tell it to listen on port 4001
$ ./p2pd node listen test /ip4/0.0.0.0/tcp/4001
```

## Contribute

PRs are welcome! Feel free to join in. All welcome. Open an [issue](https://github.com/paralin/go-libp2pd/issues)!

Check out our [contributing document](https://github.com/libp2p/community/blob/master/CONTRIBUTE.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to libp2p are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

[MIT](LICENSE) © Christian Stewart
