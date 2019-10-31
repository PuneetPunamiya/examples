// Code generated by goa v3.0.7, DO NOT EDIT.
//
// chatter gRPC server types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package server

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
)

// NewLoginPayload builds the payload of the "login" endpoint of the "chatter"
// service from the gRPC request type.
func NewLoginPayload(user string, password string) *chatter.LoginPayload {
	v := &chatter.LoginPayload{}
	v.User = user
	v.Password = password
	return v
}

// NewLoginResponse builds the gRPC response type from the result of the
// "login" endpoint of the "chatter" service.
func NewLoginResponse(result string) *chatterpb.LoginResponse {
	message := &chatterpb.LoginResponse{}
	message.Field = result
	return message
}

// NewEchoerPayload builds the payload of the "echoer" endpoint of the
// "chatter" service from the gRPC request type.
func NewEchoerPayload(token string) *chatter.EchoerPayload {
	v := &chatter.EchoerPayload{}
	v.Token = token
	return v
}

func NewEchoerResponse(result string) *chatterpb.EchoerResponse {
	v := &chatterpb.EchoerResponse{}
	v.Field = result
	return v
}

func NewEchoerStreamingRequest(v *chatterpb.EchoerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewListenerPayload builds the payload of the "listener" endpoint of the
// "chatter" service from the gRPC request type.
func NewListenerPayload(token string) *chatter.ListenerPayload {
	v := &chatter.ListenerPayload{}
	v.Token = token
	return v
}

func NewListenerStreamingRequest(v *chatterpb.ListenerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSummaryPayload builds the payload of the "summary" endpoint of the
// "chatter" service from the gRPC request type.
func NewSummaryPayload(token string) *chatter.SummaryPayload {
	v := &chatter.SummaryPayload{}
	v.Token = token
	return v
}

func NewChatSummaryCollection(vresult chatterviews.ChatSummaryCollectionView) *chatterpb.ChatSummaryCollection {
	v := &chatterpb.ChatSummaryCollection{}
	v.Field = make([]*chatterpb.ChatSummary, len(vresult))
	for i, val := range vresult {
		v.Field[i] = &chatterpb.ChatSummary{}
		if val.Message != nil {
			v.Field[i].Message_ = *val.Message
		}
		if val.Length != nil {
			v.Field[i].Length = int32(*val.Length)
		}
		if val.SentAt != nil {
			v.Field[i].SentAt = *val.SentAt
		}
	}
	return v
}

func NewSummaryStreamingRequest(v *chatterpb.SummaryStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSubscribePayload builds the payload of the "subscribe" endpoint of the
// "chatter" service from the gRPC request type.
func NewSubscribePayload(token string) *chatter.SubscribePayload {
	v := &chatter.SubscribePayload{}
	v.Token = token
	return v
}

func NewSubscribeResponse(result *chatter.Event) *chatterpb.SubscribeResponse {
	v := &chatterpb.SubscribeResponse{
		Message_: result.Message,
		Action:   result.Action,
		AddedAt:  result.AddedAt,
	}
	return v
}

// NewHistoryPayload builds the payload of the "history" endpoint of the
// "chatter" service from the gRPC request type.
func NewHistoryPayload(view *string, token string) *chatter.HistoryPayload {
	v := &chatter.HistoryPayload{}
	v.View = view
	v.Token = token
	return v
}

func NewHistoryResponse(vresult *chatterviews.ChatSummaryView) *chatterpb.HistoryResponse {
	v := &chatterpb.HistoryResponse{}
	if vresult.Message != nil {
		v.Message_ = *vresult.Message
	}
	if vresult.Length != nil {
		v.Length = int32(*vresult.Length)
	}
	if vresult.SentAt != nil {
		v.SentAt = *vresult.SentAt
	}
	return v
}
