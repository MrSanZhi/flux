package streams

import (
	"bufio"
	"fmt"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
	"io"
	"strings"
)

func init() {
	flux.RegisterPackageValue("streams", "lines", values.NewFunction(
		"lines",
		semantic.NewFunctionPolyType(semantic.FunctionPolySignature{
			Parameters: map[string]semantic.PolyType{
				"stream": semantic.Stream,
			},
			Required:     semantic.LabelSet{"stream"},
			Return:       semantic.NewArrayPolyType(semantic.String),
			PipeArgument: "stream",
		}),
		func(args values.Object) (values.Value, error) {
			stream, ok := args.Get("stream")

			if ok && stream.Type().Nature() == semantic.Stream {
				s, _ := stream.(values.Stream)
				scanner := bufio.NewScanner(s.(io.Reader))
				var lines []values.Value
				for scanner.Scan() {
					line := scanner.Text()
					//fmt.Println(line), Println will add back the final '\n'
					lines = append(lines, values.NewString(line))
				}
				if err := scanner.Err(); err != nil {
					return nil, err
				}
				return values.NewArrayWithBacking(semantic.String, lines), nil
			}
			return nil, errors.New(codes.Invalid, "argument was not a stream type")
		},
		false,
	))

	flux.RegisterPackageValue("streams", "write", values.NewFunction(
		"write",
		semantic.NewFunctionPolyType(semantic.FunctionPolySignature{
			Parameters: map[string]semantic.PolyType{
				"data": semantic.Stream,
				"to":   semantic.Stream,
			},
			Required:     semantic.LabelSet{"data", "to"},
			Return:       semantic.Bool,
			PipeArgument: "data",
		}),
		func(args values.Object) (values.Value, error) {
			to, toOk := args.Get("to")
			data, dataOk := args.Get("data")

			if toOk && dataOk && to.Type().Nature() == semantic.Stream && data.Type().Nature() == semantic.Stream {
				d, ok := data.(io.Reader)
				if !ok {
					return nil, fmt.Errorf("data stream not of type reader")
				}
				t, ok := to.(io.Writer)
				if !ok {
					return nil, fmt.Errorf("destination stream not of type writer")
				}
				_, err := io.Copy(t, d)
				if err != nil {
					return values.NewBool(false), err
				}

				return values.NewBool(true), nil
			}
			return values.NewBool(false), errors.New(codes.Invalid, "arguments are incorrect: given arguments \"to\" is of "+
				"type %t and \"data\" is of type %t", to.Type(), data.Type())
		},
		true,
	))
	flux.RegisterPackageValue("streams", "stringStream", values.NewFunction(
		"stringStream",
		semantic.NewFunctionPolyType(semantic.FunctionPolySignature{
			Parameters: map[string]semantic.PolyType{
				"v": semantic.String,
			},
			Required:     semantic.LabelSet{"v"},
			Return:       semantic.Stream,
			PipeArgument: "v",
		}),
		func(args values.Object) (values.Value, error) {
			v, ok := args.Get("v")
			if ok && v.Type().Nature() == semantic.String {
				s := strings.NewReader(v.Str())
				return values.NewReadStream(s), nil
			}
			return nil, errors.New(codes.Invalid, "cannot convert given string to stream")
		},
		false,
	))
	flux.RegisterPackageValue("streams", "zip", values.NewFunction(
		"zip",
		semantic.NewFunctionPolyType(semantic.FunctionPolySignature{
			Parameters: map[string]semantic.PolyType{
				"s": semantic.Stream,
			},
			Required:     semantic.LabelSet{"s"},
			Return:       semantic.Stream,
			PipeArgument: "s",
		}),
		func(args values.Object) (values.Value, error) {
			s, ok := args.Get("s")
			if ok && s.Type().Nature() == semantic.Stream {
				sr := s.(io.Writer)
				w := zlib.NewWriter(sr)
				return values.NewWriteStream(w), nil
			}
			return nil, errors.New(codes.Invalid, "cannot compress stream")
		},
		false,
	))
	flux.RegisterPackageValue("streams", "unzip", values.NewFunction(
		"unzip",
		semantic.NewFunctionPolyType(semantic.FunctionPolySignature{
			Parameters: map[string]semantic.PolyType{
				"s": semantic.Stream,
			},
			Required:     semantic.LabelSet{"s"},
			Return:       semantic.Stream,
			PipeArgument: "s",
		}),
		func(args values.Object) (values.Value, error) {
			s, ok := args.Get("s")
			if ok && s.Type().Nature() == semantic.Stream {
				r, _ := zlib.NewReader(s.(io.Reader))
				return values.NewReadStream(r), nil
			}
			return nil, errors.New(codes.Invalid, "cannot decompress stream")
		},
		false,
	))
}
