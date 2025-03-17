# cyberdeck

Various homecooked cyber security tools.

> [!WARNING]
> Use at your own risk!

Contents:
- [tarpit](./README.md#tarpit): stall unwanted network connections

# tarpit
Based on [endlessh](https://github.com/skeeto/endlessh).
```sh
go run ssh-tarpit/main.go
```
TODO:
- use iptables to drop/stall connections to save on resources (aka blackhole option)
- enable tarpitting of other/multiple protocols & ports
