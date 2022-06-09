package reader

import "io"

var _ io.ReadCloser = readCloser{}

type readCloser struct {
	reader io.Reader
	closer io.Closer
}

func WithCloser(reader io.Reader, closer io.Closer) io.ReadCloser {
	return readCloser{
		reader: reader,
		closer: closer,
	}
}

// Read implements io.ReadCloser.
func (rc readCloser) Read(p []byte) (int, error) {
	return rc.reader.Read(p)
}

// Close implements io.ReadCloser.
func (rc readCloser) Close() error {
	return rc.closer.Close()
}
