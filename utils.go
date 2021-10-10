package opfr

import (
	"bytes"
	"math/big"
)

// https://tools.ietf.org/html/rfc8017#section-4

// i2osp converts an integer to an octet-string
// (https://tools.ietf.org/html/rfc8017#section-4.1)
func i2ospFromInt(x, xLen int) ([]byte, error) {
	if x < 0 {
		return nil, ErrI2OSPNegativeInteger
	}

	if x >= (1 << (8 * xLen)) {
		return nil, ErrI2OSPLargeInteger
	}

	ret := make([]byte, xLen)
	val := x
	for i := xLen - 1; i >= 0; i-- {
		ret[i] = byte(val & 0xff)
		val = val >> 8
	}

	return ret, nil
}

func i2osp(b *big.Int, n int) ([]byte, error) {
	os := b.Bytes()
	if n > len(os) {
		var buf bytes.Buffer
		buf.Write(make([]byte, n-len(os))) // prepend 0s
		buf.Write(os)
		return buf.Bytes()
	}

	return os[:n]
}

// os2ip converts an octet-string to an integer
// (https://tools.ietf.org/html/rfc8017#section-4.2)
func os2ip(x []byte) *big.Int {
	return new(big.Int).SetBytes(x)
}
