// Code generated by goa v3.0.7, DO NOT EDIT.
//
// text HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	text "goa.design/examples/encodings/gen/text"
	goahttp "goa.design/goa/v3/http"
)

// BuildConcatstringsRequest instantiates a HTTP request object with method and
// path set to call the "text" service "concatstrings" endpoint
func (c *Client) BuildConcatstringsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a string
		b string
	)
	{
		p, ok := v.(*text.ConcatstringsPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("text", "concatstrings", "*text.ConcatstringsPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConcatstringsTextPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("text", "concatstrings", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConcatstringsResponse returns a decoder for responses returned by the
// text concatstrings endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeConcatstringsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("text", "concatstrings", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("text", "concatstrings", resp.StatusCode, string(body))
		}
	}
}

// BuildConcatbytesRequest instantiates a HTTP request object with method and
// path set to call the "text" service "concatbytes" endpoint
func (c *Client) BuildConcatbytesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a string
		b string
	)
	{
		p, ok := v.(*text.ConcatbytesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("text", "concatbytes", "*text.ConcatbytesPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConcatbytesTextPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("text", "concatbytes", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConcatbytesResponse returns a decoder for responses returned by the
// text concatbytes endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeConcatbytesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []byte
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("text", "concatbytes", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("text", "concatbytes", resp.StatusCode, string(body))
		}
	}
}

// BuildConcatstringfieldRequest instantiates a HTTP request object with method
// and path set to call the "text" service "concatstringfield" endpoint
func (c *Client) BuildConcatstringfieldRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a string
		b string
	)
	{
		p, ok := v.(*text.ConcatstringfieldPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("text", "concatstringfield", "*text.ConcatstringfieldPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConcatstringfieldTextPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("text", "concatstringfield", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConcatstringfieldResponse returns a decoder for responses returned by
// the text concatstringfield endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeConcatstringfieldResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("text", "concatstringfield", err)
			}
			res := NewConcatstringfieldMyConcatenationOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("text", "concatstringfield", resp.StatusCode, string(body))
		}
	}
}

// BuildConcatbytesfieldRequest instantiates a HTTP request object with method
// and path set to call the "text" service "concatbytesfield" endpoint
func (c *Client) BuildConcatbytesfieldRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a string
		b string
	)
	{
		p, ok := v.(*text.ConcatbytesfieldPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("text", "concatbytesfield", "*text.ConcatbytesfieldPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ConcatbytesfieldTextPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("text", "concatbytesfield", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeConcatbytesfieldResponse returns a decoder for responses returned by
// the text concatbytesfield endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeConcatbytesfieldResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []byte
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("text", "concatbytesfield", err)
			}
			res := NewConcatbytesfieldMyConcatenationOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("text", "concatbytesfield", resp.StatusCode, string(body))
		}
	}
}
