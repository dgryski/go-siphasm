// Package siphasm is an assembly implementation of siphash-2-4
/*

It is a translation of the assembly produced by GCC 4.8.2 for the pure-C siphash implementation by Floodyberry ( https://github.com/floodyberry/siphash )

It has a 1.3x-1.8x speedup over the pure-Go implementation ( https://github.com/dchest/siphash )

To the extent possible, this is donated to the public domain.

*/
package siphasm

// Hash returns the 64-bit SipHash-2-4 of the given byte slice with two 64-bit
// parts of 128-bit key: k0 and k1.
func Hash(k0, k1 uint64, b []byte) uint64
