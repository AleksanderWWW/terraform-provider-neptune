package neptune

import (
	"fmt"
	"io"
	"net/http"
)

// MockReadCloser is a mock implementation of io.ReadCloser.
type MockReadCloser struct {
	data  string
	index int
}

// NewMockReadCloser creates a new instance of MockReadCloser.
func NewMockReadCloser(data string) *MockReadCloser {
	return &MockReadCloser{data: data}
}

// Read implements the io.Reader interface.
func (m *MockReadCloser) Read(p []byte) (n int, err error) {
	if m.index >= len(m.data) {
		return 0, io.EOF
	}

	n = copy(p, m.data[m.index:])
	m.index += n
	return n, nil
}

// Close implements the io.Closer interface.
func (m *MockReadCloser) Close() error {
	return nil
}

type MockHttpClient struct {
	resps   []*http.Response
	errs    []error
	reqSent *http.Request
	index   int
}

func (mhc *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	mhc.reqSent = req

	if mhc.index >= len(mhc.resps) {
		resp = &http.Response{}
	} else {
		resp = mhc.resps[mhc.index]
	}

	if mhc.index >= len(mhc.errs) {
		err = fmt.Errorf("Too few errors provided for mock http client")
	} else {
		err = mhc.errs[mhc.index]
	}

	mhc.index++
	return resp, err
}
