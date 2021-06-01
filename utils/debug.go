package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func CopyBody(req *http.Request) ([]byte, error) {
	var err error
	save := req.Body
	if req.Body == nil {
		req.Body = nil
	} else {
		save, req.Body, err = drainBody(req.Body)
		if err != nil {
			return nil, err
		}
	}

	var b bytes.Buffer
	if req.Body != nil {
		var dest io.Writer = &b
		_, err = io.Copy(dest, req.Body)
		if err != nil {
			return nil, err
		}
	}

	req.Body = save
	return b.Bytes(), nil
}
