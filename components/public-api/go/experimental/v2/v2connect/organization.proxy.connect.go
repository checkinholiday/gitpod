// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v2connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v2 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v2"
)

var _ OrganizationServiceHandler = (*ProxyOrganizationServiceHandler)(nil)

type ProxyOrganizationServiceHandler struct {
	Client v2.OrganizationServiceClient
	UnimplementedOrganizationServiceHandler
}

func (s *ProxyOrganizationServiceHandler) CreateOrganization(ctx context.Context, req *connect_go.Request[v2.CreateOrganizationRequest]) (*connect_go.Response[v2.CreateOrganizationResponse], error) {
	resp, err := s.Client.CreateOrganization(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) GetOrganization(ctx context.Context, req *connect_go.Request[v2.GetOrganizationRequest]) (*connect_go.Response[v2.GetOrganizationResponse], error) {
	resp, err := s.Client.GetOrganization(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) UpdateOrganization(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationRequest]) (*connect_go.Response[v2.UpdateOrganizationResponse], error) {
	resp, err := s.Client.UpdateOrganization(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) ListOrganizations(ctx context.Context, req *connect_go.Request[v2.ListOrganizationsRequest]) (*connect_go.Response[v2.ListOrganizationsResponse], error) {
	resp, err := s.Client.ListOrganizations(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) DeleteOrganization(ctx context.Context, req *connect_go.Request[v2.DeleteOrganizationRequest]) (*connect_go.Response[v2.DeleteOrganizationResponse], error) {
	resp, err := s.Client.DeleteOrganization(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) GetOrganizationInvitation(ctx context.Context, req *connect_go.Request[v2.GetOrganizationInvitationRequest]) (*connect_go.Response[v2.GetOrganizationInvitationResponse], error) {
	resp, err := s.Client.GetOrganizationInvitation(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) JoinOrganization(ctx context.Context, req *connect_go.Request[v2.JoinOrganizationRequest]) (*connect_go.Response[v2.JoinOrganizationResponse], error) {
	resp, err := s.Client.JoinOrganization(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) ResetOrganizationInvitation(ctx context.Context, req *connect_go.Request[v2.ResetOrganizationInvitationRequest]) (*connect_go.Response[v2.ResetOrganizationInvitationResponse], error) {
	resp, err := s.Client.ResetOrganizationInvitation(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) ListOrganizationMembers(ctx context.Context, req *connect_go.Request[v2.ListOrganizationMembersRequest]) (*connect_go.Response[v2.ListOrganizationMembersResponse], error) {
	resp, err := s.Client.ListOrganizationMembers(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) UpdateOrganizationMember(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationMemberRequest]) (*connect_go.Response[v2.UpdateOrganizationMemberResponse], error) {
	resp, err := s.Client.UpdateOrganizationMember(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) DeleteOrganizationMember(ctx context.Context, req *connect_go.Request[v2.DeleteOrganizationMemberRequest]) (*connect_go.Response[v2.DeleteOrganizationMemberResponse], error) {
	resp, err := s.Client.DeleteOrganizationMember(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) GetOrganizationSettings(ctx context.Context, req *connect_go.Request[v2.GetOrganizationSettingsRequest]) (*connect_go.Response[v2.GetOrganizationSettingsResponse], error) {
	resp, err := s.Client.GetOrganizationSettings(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOrganizationServiceHandler) UpdateOrganizationSettings(ctx context.Context, req *connect_go.Request[v2.UpdateOrganizationSettingsRequest]) (*connect_go.Response[v2.UpdateOrganizationSettingsResponse], error) {
	resp, err := s.Client.UpdateOrganizationSettings(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}