package writer

import "io"

var _ io.WriteCloser = writeCloser{}

type writeCloser struct {
	writer io.Writer
	closer io.Closer
}

func WithCloser(writer io.Writer, closer io.Closer) io.WriteCloser {
	return writeCloser{
		writer: writer,
		closer: closer,
	}
}

// Write implements io.WriteCloser.
func (wc writeCloser) Write(p []byte) (int, error) {
	return wc.writer.Write(p)
}

// Close implements io.WriteCloser.
func (wc writeCloser) Close() error {
	return wc.closer.Close()
}
