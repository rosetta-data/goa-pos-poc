// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pos": Purchase Resource Client
//
// Command:
// $ goagen
// --design=github.com/psavelis/goa-pos-poc/design
// --out=$(GOPATH)src\github.com\psavelis\goa-pos-poc
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreatePurchasePath computes a request path to the create action of Purchase.
func CreatePurchasePath() string {

	return fmt.Sprintf("/purchases/")
}

// creates a purchase
func (c *Client) CreatePurchase(ctx context.Context, path string, payload *PurchasePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePurchaseRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePurchaseRequest create the request corresponding to the create action endpoint of the Purchase resource.
func (c *Client) NewCreatePurchaseRequest(ctx context.Context, path string, payload *PurchasePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// FindPurchasePath computes a request path to the find action of Purchase.
func FindPurchasePath(transactionID string) string {
	param0 := transactionID

	return fmt.Sprintf("/purchases/%s", param0)
}

// retrieve an specific purchase
func (c *Client) FindPurchase(ctx context.Context, path string, payload *PurchasePayload, contentType string) (*http.Response, error) {
	req, err := c.NewFindPurchaseRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFindPurchaseRequest create the request corresponding to the find action endpoint of the Purchase resource.
func (c *Client) NewFindPurchaseRequest(ctx context.Context, path string, payload *PurchasePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
