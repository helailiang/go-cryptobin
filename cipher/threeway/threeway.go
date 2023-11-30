package threeway

import (
    "unsafe"
    "strconv"
    "crypto/cipher"
    "encoding/binary"
)

const BlockSize = 12

type threewayCipher struct {
    key [3]uint32
}

// NewCipher creates and returns a new cipher.Block.
// data bytes use BigEndian
func NewCipher(key []byte) (cipher.Block, error) {
    k := len(key)
    switch k {
        case 12:
            break
        default:
            return nil, KeySizeError(len(key))
    }

    c := new(threewayCipher)
    c.key[0] = binary.BigEndian.Uint32(key[0:])
    c.key[1] = binary.BigEndian.Uint32(key[4:])
    c.key[2] = binary.BigEndian.Uint32(key[8:])

    return c, nil
}

func (this *threewayCipher) BlockSize() int {
    return BlockSize
}

func (this *threewayCipher) Encrypt(dst, src []byte) {
    if len(src) < BlockSize {
        panic("crypto/loki97: input not full block")
    }

    if len(dst) < BlockSize {
        panic("crypto/loki97: output not full block")
    }

    if inexactOverlap(dst[:BlockSize], src[:BlockSize]) {
        panic("crypto/loki97: invalid buffer overlap")
    }

    this.encrypt(dst, src)
}

func (this *threewayCipher) Decrypt(dst, src []byte) {
    if len(src) < BlockSize {
        panic("crypto/loki97: input not full block")
    }

    if len(dst) < BlockSize {
        panic("crypto/loki97: output not full block")
    }

    if inexactOverlap(dst[:BlockSize], src[:BlockSize]) {
        panic("crypto/loki97: invalid buffer overlap")
    }

    this.decrypt(dst, src)
}

func (this *threewayCipher) encrypt(dst, src []byte) {
    var rcon [12]uint32

    var ciphertext [3]uint32

    ciphertext[0] = binary.BigEndian.Uint32(src[0:])
    ciphertext[1] = binary.BigEndian.Uint32(src[4:])
    ciphertext[2] = binary.BigEndian.Uint32(src[8:])

    rndcon_gen(STRT_E, &rcon)

    var i uint32

    for i = 0; i < NMBR; i++ {
        ciphertext[0] ^= this.key[0] ^ (rcon[i] << 16)
        ciphertext[1] ^= this.key[1]
        ciphertext[2] ^= this.key[2] ^ rcon[i]
        rho(&ciphertext)
    }

    ciphertext[0] ^= this.key[0] ^ (rcon[NMBR] << 16)
    ciphertext[1] ^= this.key[1]
    ciphertext[2] ^= this.key[2] ^ rcon[NMBR]

    theta(&ciphertext)

    var data [12]byte

    binary.BigEndian.PutUint32(data[0:], ciphertext[0])
    binary.BigEndian.PutUint32(data[4:], ciphertext[1])
    binary.BigEndian.PutUint32(data[8:], ciphertext[2])

    copy(dst, data[:])
}

func (this *threewayCipher) decrypt(dst, src []byte) {
    var rcon [12]uint32 // the `inverse' round constants

    var plaintext [3]uint32

    plaintext[0] = binary.BigEndian.Uint32(src[0:])
    plaintext[1] = binary.BigEndian.Uint32(src[4:])
    plaintext[2] = binary.BigEndian.Uint32(src[8:])

    var key [3]uint32

    for i, k := range this.key {
        key[i] = k
    }

    theta(&key)
    mu(&key)

    rndcon_gen(STRT_D, &rcon)

    mu(&plaintext)

    var i uint32

    for i = 0; i < NMBR; i++ {
        plaintext[0] ^= key[0] ^ (rcon[i] << 16)
        plaintext[1] ^= key[1]
        plaintext[2] ^= key[2] ^ rcon[i]
        rho(&plaintext)
    }

    plaintext[0] ^= key[0] ^ (rcon[NMBR] << 16)
    plaintext[1] ^= key[1]
    plaintext[2] ^= key[2] ^ rcon[NMBR]

    theta(&plaintext)
    mu(&plaintext)

    var data [12]byte

    binary.BigEndian.PutUint32(data[0:], plaintext[0])
    binary.BigEndian.PutUint32(data[4:], plaintext[1])
    binary.BigEndian.PutUint32(data[8:], plaintext[2])

    copy(dst, data[:])
}

// anyOverlap reports whether x and y share memory at any (not necessarily
// corresponding) index. The memory beyond the slice length is ignored.
func anyOverlap(x, y []byte) bool {
    return len(x) > 0 && len(y) > 0 &&
        uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
        uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}

// inexactOverlap reports whether x and y share memory at any non-corresponding
// index. The memory beyond the slice length is ignored. Note that x and y can
// have different lengths and still not have any inexact overlap.
//
// inexactOverlap can be used to implement the requirements of the crypto/cipher
// AEAD, Block, BlockMode and Stream interfaces.
func inexactOverlap(x, y []byte) bool {
    if len(x) == 0 || len(y) == 0 || &x[0] == &y[0] {
        return false
    }

    return anyOverlap(x, y)
}

type KeySizeError int

func (k KeySizeError) Error() string {
    return "crypto/threeway: invalid key size " + strconv.Itoa(int(k))
}
