package closer

import "io"

var _ io.Closer = Func(nil)

type Func func() error

// Close implements io.Closer.
func (f Func) Close() error {
	return f()
}
