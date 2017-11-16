// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "My API": Operands Resource Client
//
// Command:
// $ goagen
// --design=github.com/linh-snoopy/go-examples/goatest/design
// --out=c:\Users\LENOVO\go\src\github.com\linh-snoopy\go-examples\goatest\gen
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// SumOperandsPath computes a request path to the sum action of Operands.
func SumOperandsPath(left int, right int) string {
	param0 := strconv.Itoa(left)
	param1 := strconv.Itoa(right)

	return fmt.Sprintf("/results/sum/%s/%s", param0, param1)
}

// Sum
func (c *Client) SumOperands(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSumOperandsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSumOperandsRequest create the request corresponding to the sum action endpoint of the Operands resource.
func (c *Client) NewSumOperandsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
