package msgpack_test

import (
	"bytes"
	"encoding/hex"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/KyberNetwork/msgpack/v5"
	"github.com/davecgh/go-spew/spew"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

type Address [20]byte

func hexToAddress(s string) Address {
	s = strings.TrimPrefix(s, "0x")
	var a Address
	b, err := hex.DecodeString(s)
	if err != nil {
		return a
	}
	if len(b) > 20 {
		copy(a[:], b[len(b)-20:])
	} else {
		copy(a[20-len(b):], b[:])
	}
	return a
}

type Bar struct {
	alpha *uint256.Int
}

type IBar interface {
	Bar()
}

func (*Bar) Bar() {}

type tokenInfo struct {
	indexPlus1 uint8
	scale      uint8
	gauge      Address
}

type Foo struct {
	alpha   *uint64
	beta    *big.Int
	gamma   *uint256.Int
	delta   uint256.Int
	epsilon IBar
	zeta    map[uint64]string
	eta     []string
	theta   map[string]string
	iota_   map[string]bool
	kappa   map[string]struct{}
	lambda  map[string]tokenInfo
	mu      *Token
}

func newUint64(v uint64) *uint64 { return &v }

type Currency interface {
	Currency()
}

type BaseCurrency struct {
	Currency Currency `msgpack:"-"`
	Symbol   string
	Name     string
}

type Token struct {
	*BaseCurrency
	Address string
}

func (*Token) Currency() {}

func (t *Token) AfterMsgpackUnmarshal() error {
	t.BaseCurrency.Currency = t
	return nil
}

func TestMarshalUnmarshalForked(t *testing.T) {
	msgpack.RegisterConcreteType(&Bar{})

	expected := &Foo{
		alpha: newUint64(9999999999999999),
		beta:  new(big.Int).Exp(new(big.Int).SetUint64(math.MaxUint64), big.NewInt(4), nil),
		gamma: new(uint256.Int).Exp(new(uint256.Int).SetUint64(math.MaxUint64), uint256.NewInt(4)),
		delta: *new(uint256.Int).Exp(new(uint256.Int).SetUint64(math.MaxUint64), uint256.NewInt(4)),
		epsilon: &Bar{
			alpha: new(uint256.Int).Exp(new(uint256.Int).SetUint64(math.MaxUint64), uint256.NewInt(4)),
		},
		zeta: map[uint64]string{
			69696969: "foo",
		},
		eta: []string{"foo", "bar"},
		theta: map[string]string{
			"foo": "bar",
			"abc": "def",
		},
		iota_: map[string]bool{
			"foo": true,
			"abc": false,
		},
		kappa: map[string]struct{}{
			"foo": {},
			"bar": {},
		},
		lambda: map[string]tokenInfo{
			"foo": {
				indexPlus1: 99,
				scale:      99,
				gauge:      hexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			},
		},
		mu: &Token{
			BaseCurrency: &BaseCurrency{
				Symbol: "T1",
				Name:   "Token 1",
			},
			Address: "0xaaaa",
		},
	}
	expected.mu.BaseCurrency.Currency = expected.mu

	var encoded bytes.Buffer
	en := msgpack.NewEncoder(&encoded)
	en.IncludeUnexported(true)
	en.SetForceAsArray(true)

	err := en.Encode(expected)
	require.NoError(t, err)

	decoded := new(Foo)
	de := msgpack.NewDecoder(&encoded)
	de.IncludeUnexported(true)
	en.SetForceAsArray(true)

	err = de.Decode(decoded)
	require.NoError(t, err)

	spew.Dump(decoded)

	require.EqualValues(t, expected, decoded)
}
