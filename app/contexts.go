// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "pos": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/psavelis/goa-pos-poc/design
// --out=$(GOPATH)src\github.com\psavelis\goa-pos-poc
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreatePurchaseContext provides the Purchase create action context.
type CreatePurchaseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PurchasePayload
}

// NewCreatePurchaseContext parses the incoming request URL and body, performs validations and creates the
// context used by the Purchase controller create action.
func NewCreatePurchaseContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePurchaseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePurchaseContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePurchaseContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreatePurchaseContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// Conflict sends a HTTP response with status code 409.
func (ctx *CreatePurchaseContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// ShowPurchaseContext provides the Purchase show action context.
type ShowPurchaseContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TransactionID string
}

// NewShowPurchaseContext parses the incoming request URL and body, performs validations and creates the
// context used by the Purchase controller show action.
func NewShowPurchaseContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowPurchaseContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowPurchaseContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTransactionID := req.Params["transaction_id"]
	if len(paramTransactionID) > 0 {
		rawTransactionID := paramTransactionID[0]
		rctx.TransactionID = rawTransactionID
		if ok := goa.ValidatePattern(`^[0-9a-fA-F]{24}$`, rctx.TransactionID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`transaction_id`, rctx.TransactionID, `^[0-9a-fA-F]{24}$`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPurchaseContext) OK(r *Purchase) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowPurchaseContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPurchaseContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
