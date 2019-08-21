// Code generated by goa v2.0.4, DO NOT EDIT.
//
// storage gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package server

import (
	"context"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeListResponse encodes responses from the "storage" service "list"
// endpoint.
func EncodeListResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(storageviews.StoredBottleCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "list", "storageviews.StoredBottleCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewStoredBottleCollection(result)
	return resp, nil
}

// EncodeShowResponse encodes responses from the "storage" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*storageviews.StoredBottle)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "show", "*storageviews.StoredBottle", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "storage" service "show" endpoint.
func DecodeShowRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view *string
		err  error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *storagepb.ShowRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.ShowRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "show", "*storagepb.ShowRequest", v)
		}
	}
	var payload *storage.ShowPayload
	{
		payload = NewShowPayload(message, view)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "storage" service "add"
// endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "add", "string", v)
	}
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "storage" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *storagepb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "add", "*storagepb.AddRequest", v)
		}
		if err := ValidateAddRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *storage.Bottle
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeRemoveResponse encodes responses from the "storage" service "remove"
// endpoint.
func EncodeRemoveResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewRemoveResponse()
	return resp, nil
}

// DecodeRemoveRequest decodes requests sent to "storage" service "remove"
// endpoint.
func DecodeRemoveRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *storagepb.RemoveRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.RemoveRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "remove", "*storagepb.RemoveRequest", v)
		}
	}
	var payload *storage.RemovePayload
	{
		payload = NewRemovePayload(message)
	}
	return payload, nil
}

// EncodeRateResponse encodes responses from the "storage" service "rate"
// endpoint.
func EncodeRateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewRateResponse()
	return resp, nil
}

// DecodeRateRequest decodes requests sent to "storage" service "rate" endpoint.
func DecodeRateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *storagepb.RateRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.RateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "rate", "*storagepb.RateRequest", v)
		}
		if err := ValidateRateRequest(message); err != nil {
			return nil, err
		}
	}
	var payload map[uint32][]string
	{
		payload = NewRatePayload(message)
	}
	return payload, nil
}

// EncodeMultiAddResponse encodes responses from the "storage" service
// "multi_add" endpoint.
func EncodeMultiAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.([]string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("storage", "multi_add", "[]string", v)
	}
	resp := NewMultiAddResponse(result)
	return resp, nil
}

// DecodeMultiAddRequest decodes requests sent to "storage" service "multi_add"
// endpoint.
func DecodeMultiAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *storagepb.MultiAddRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.MultiAddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "multi_add", "*storagepb.MultiAddRequest", v)
		}
		if err := ValidateMultiAddRequest(message); err != nil {
			return nil, err
		}
	}
	var payload []*storage.Bottle
	{
		payload = NewMultiAddPayload(message)
	}
	return payload, nil
}

// EncodeMultiUpdateResponse encodes responses from the "storage" service
// "multi_update" endpoint.
func EncodeMultiUpdateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewMultiUpdateResponse()
	return resp, nil
}

// DecodeMultiUpdateRequest decodes requests sent to "storage" service
// "multi_update" endpoint.
func DecodeMultiUpdateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *storagepb.MultiUpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*storagepb.MultiUpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("storage", "multi_update", "*storagepb.MultiUpdateRequest", v)
		}
		if err := ValidateMultiUpdateRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *storage.MultiUpdatePayload
	{
		payload = NewMultiUpdatePayload(message)
	}
	return payload, nil
}
