package main

import (
	"context"
	basic "github.com/Redmoon-2333/innospark-idl-personal/kitex_gen/basic"
	core_api "github.com/Redmoon-2333/innospark-idl-personal/kitex_gen/core_api"
)

// CoreApiImpl implements the last service interface defined in the IDL.
type CoreApiImpl struct{}

func (s *CoreApiImpl) Completions(req *core_api.CompletionsReq, stream core_api.CoreApi_CompletionsServer) (err error) {
	println("Completions called")
	return
}

// ListConversation implements the CoreApiImpl interface.
func (s *CoreApiImpl) ListConversation(ctx context.Context, req *core_api.ListConversationReq) (resp *core_api.ListConversationResp, err error) {
	// TODO: Your code here...
	return
}

// GetConversation implements the CoreApiImpl interface.
func (s *CoreApiImpl) GetConversation(ctx context.Context, req *core_api.GetConversationReq) (resp *core_api.GetConversationResp, err error) {
	// TODO: Your code here...
	return
}

// ListAgents implements the CoreApiImpl interface.
func (s *CoreApiImpl) ListAgents(ctx context.Context, req *core_api.ListAgentsReq) (resp *core_api.ListAgentsResp, err error) {
	// TODO: Your code here...
	return
}

// Feedback implements the CoreApiImpl interface.
func (s *CoreApiImpl) Feedback(ctx context.Context, req *core_api.FeedbackReq) (resp *basic.Response, err error) {
	// TODO: Your code here...
	return
}
