package opensearch

import (
	"io"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

// BuildRequest is a helper function to build a http.Request
func BuildRequest(method string, path string, body io.Reader, params map[string]string, headers http.Header) (*http.Request, error) {
	//nolint:noctx // ctx gets appended when the requests gets executet
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := httpReq.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}

		httpReq.URL.RawQuery = q.Encode()
	}

	if body != nil {
		httpReq.Header[headerContentType] = []string{"application/json"}
	}

	if len(headers) > 0 {
		if len(httpReq.Header) == 0 {
			httpReq.Header = headers
		} else {
			for k, vv := range headers {
				for _, v := range vv {
					httpReq.Header.Add(k, v)
				}
			}
		}
	}
	return httpReq, nil
}
