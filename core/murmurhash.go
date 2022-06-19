package core

func ROTL32(x uint32, r int8) uint32 {
	return (x << r) | (x >> (32 - r))
}

/* murmurHash3 x86_32bits implementation of murmurHash3 */
func MurmurHash3(value []byte, length int, seed uint32) uint32 {
	/* variables */
	nblocks := length / 4
	h1 := seed
	data := value
	/* constants */
	const c1 uint32 = 0xcc9e2d51
	const c2 uint32 = 0x1b873593

	/* body */
	var k1 uint32
	for i := 0; i < nblocks; i++ {
		k1 = uint32(data[i*4+0]) + uint32(data[i*4+1])<<8 + uint32(data[i*4+2])<<16 + uint32(data[i*4+3])<<24

		k1 *= c1
		k1 = ROTL32(k1, 15)
		k1 *= c2

		h1 ^= k1
		h1 = ROTL32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	/* tail */
	tail := data[nblocks*4:]

	switch length & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = ROTL32(k1, 15)
		k1 *= c2
		h1 ^= k1
	}

	/* finalization */
	h1 ^= uint32(length)
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}
