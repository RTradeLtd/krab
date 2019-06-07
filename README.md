# krab

`krab` is a "semi secure" keystore that satisfies the IPFS keystore interface, allowing it to be used natively with many existing IPFS implementations, and tools. It stores keys on disk in a badger datastore, encrypting the keys before being stored in the datastore. Each time a key is fetched, it is decrypted first. A single password is used to encrypt all keys.

# future improvements

* Define groups of keys, with each group having a separate encryption password