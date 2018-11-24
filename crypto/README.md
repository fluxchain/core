## Flux Crypto

Internal crypto library built on Golang's crypto


### Signatures

Flux uses [ed25519](https://ed25519.cr.yp.to/) scheme for all signatures
ed25519 is a compact schnorr scheme on Bernstein's curve 25519 .

Ed25519 has 32 byte public keys and 64 byte signatures .


### Hashes

Blake2b is used everywhere, Blake2b is length-extension resistant fast and optimized for 64 bit CPUs


