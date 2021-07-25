package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type parseCase struct {
	in  interface{}
	out interface{}
	err error
}

func TestParse(t *testing.T) {
	Convey("test int case", t, func() {
		intCases := []*parseCase{
			&parseCase{
				in:  int(123),
				out: 123,
			},
			&parseCase{
				in:  int8(123),
				out: 123,
			},
			&parseCase{
				in:  int16(123),
				out: 123,
			},
			&parseCase{
				in:  int32(123),
				out: 123,
			},
			&parseCase{
				in:  int64(123),
				out: 123,
			},
			&parseCase{
				in:  uint(123),
				out: 123,
			},
			&parseCase{
				in:  uint8(123),
				out: 123,
			},
			&parseCase{
				in:  uint16(123),
				out: 123,
			},
			&parseCase{
				in:  uint32(123),
				out: 123,
			},
			&parseCase{
				in:  uint64(123),
				out: 123,
			},
			&parseCase{
				in:  float32(123.4),
				out: 123,
			},
			&parseCase{
				in:  float64(123.5),
				out: 123,
			},
			&parseCase{
				in:  float32(123.6),
				out: 123,
			},
			&parseCase{
				in:  "123",
				out: 123,
			},
			&parseCase{
				in:  "123.4",
				out: 0,
				err: ErrConvertToIntError,
			},
			&parseCase{
				in:  "123a",
				out: 0,
				err: ErrConvertToIntError,
			},
		}

		for _, c := range intCases {
			v, err := ParseInt64(c.in)
			So(err, ShouldEqual, c.err)
			So(v, ShouldEqual, c.out)
		}
	})

	Convey("test float case", t, func() {
		intCases := []*parseCase{
			&parseCase{
				in:  int(123),
				out: 123.0,
			},
			&parseCase{
				in:  int8(123),
				out: 123.0,
			},
			&parseCase{
				in:  int16(123),
				out: 123.0,
			},
			&parseCase{
				in:  int32(123),
				out: 123.0,
			},
			&parseCase{
				in:  int64(123),
				out: 123.0,
			},
			&parseCase{
				in:  uint(123),
				out: 123.0,
			},
			&parseCase{
				in:  uint8(123),
				out: 123.0,
			},
			&parseCase{
				in:  uint16(123),
				out: 123.0,
			},
			&parseCase{
				in:  uint32(123),
				out: 123.0,
			},
			&parseCase{
				in:  uint64(123),
				out: 123.0,
			},
			&parseCase{
				in:  float32(123.4),
				out: 123.4,
			},
			&parseCase{
				in:  float64(123.5),
				out: 123.5,
			},
			&parseCase{
				in:  float32(123.6),
				out: 123.6,
			},
			&parseCase{
				in:  "123",
				out: 123.0,
			},
			&parseCase{
				in:  "123a",
				out: 0.0,
				err: ErrConvertToFloatError,
			},
		}

		for _, c := range intCases {
			v, err := ParseFloat64(c.in)
			So(err, ShouldEqual, c.err)
			So(fmt.Sprintf("%.1f", v), ShouldEqual, fmt.Sprintf("%.1f", c.out))
		}
	})
}
