package msgpack_test

import (
	"bytes"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/KyberNetwork/msgpack/v5"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

type Bar struct {
	alpha *uint256.Int
}

type IBar interface {
	Bar()
}

func (*Bar) Bar() {}

type Foo struct {
	alpha   *uint64
	beta    *big.Int
	gamma   *uint256.Int
	delta   uint256.Int
	epsilon IBar
	zeta    map[uint64]string
	ignored uint64
}

func (f *Foo) BeforeMsgpackMarshal() error {
	*f.alpha += 1
	return nil
}

func (f *Foo) AfterMsgpackUnmarshal() error {
	*f.alpha -= 1
	return nil
}

func newUint64(v uint64) *uint64 { return &v }

func TestFoo(t *testing.T) {
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
		ignored: 69696969,
	}

	var encoded bytes.Buffer
	en := msgpack.NewEncoder(&encoded)
	en.IncludeUnexported(true)
	en.IgnoreStructField(reflect.TypeOf((*Foo)(nil)).Elem(), "ignored")

	err := en.Encode(expected)
	require.NoError(t, err)

	decoded := new(Foo)
	de := msgpack.NewDecoder(&encoded)
	de.IncludeUnexported(true)
	de.IgnoreStructField(reflect.TypeOf((*Foo)(nil)).Elem(), "ignored")

	err = de.Decode(decoded)
	require.NoError(t, err)

	require.Zerof(t, decoded.ignored, "ignored fields via IgnoreStructField must be zero")
	decoded.ignored = expected.ignored

	*expected.alpha -= 1
	require.EqualValues(t, expected, decoded)
}
