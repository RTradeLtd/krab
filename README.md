# krab

`krab` is a "semi secure" keystore that satisfies the IPFS keystore interface, allowing it to be used natively with many existing IPFS implementations, and tools. It stores keys on disk in a badger datastore, encrypting the keys before being stored in the datastore. Each time a key is fetched, it is decrypted first. A single password is used to encrypt all keys.

## Multi-Language

[![](https://img.shields.io/badge/Lang-English-blue.svg)](README.md)  [![jaywcjlove/sb](https://jaywcjlove.github.io/sb/lang/chinese.svg)](README-zh.md)

## usage

There are two ways to use `krab`, one is to import with `import "github.com/RTradeLtd/krab"`, the other is to use [RTradeLtd/kaas](https://github.com/RTradeLtd/kaas) which exposes `krab` functionality via a gRPC API.

## limitations

Because badger is used as the underlying data store, a single badger datastore is unable to have multiple indepent readers/writers, only readers. If you are writing to the datastore (aka, storing keys) then you must use the gRPC version, as it enables usage of a single badger datastore via multiple different services.

## future improvements

* Define groups of keys, with each group having a separate encryption password