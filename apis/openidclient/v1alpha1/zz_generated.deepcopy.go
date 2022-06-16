//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFlowBindingOverridesObservation) DeepCopyInto(out *AuthenticationFlowBindingOverridesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFlowBindingOverridesObservation.
func (in *AuthenticationFlowBindingOverridesObservation) DeepCopy() *AuthenticationFlowBindingOverridesObservation {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFlowBindingOverridesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationFlowBindingOverridesParameters) DeepCopyInto(out *AuthenticationFlowBindingOverridesParameters) {
	*out = *in
	if in.BrowserID != nil {
		in, out := &in.BrowserID, &out.BrowserID
		*out = new(string)
		**out = **in
	}
	if in.DirectGrantID != nil {
		in, out := &in.DirectGrantID, &out.DirectGrantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationFlowBindingOverridesParameters.
func (in *AuthenticationFlowBindingOverridesParameters) DeepCopy() *AuthenticationFlowBindingOverridesParameters {
	if in == nil {
		return nil
	}
	out := new(AuthenticationFlowBindingOverridesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationObservation) DeepCopyInto(out *AuthorizationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationObservation.
func (in *AuthorizationObservation) DeepCopy() *AuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(AuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationParameters) DeepCopyInto(out *AuthorizationParameters) {
	*out = *in
	if in.AllowRemoteResourceManagement != nil {
		in, out := &in.AllowRemoteResourceManagement, &out.AllowRemoteResourceManagement
		*out = new(bool)
		**out = **in
	}
	if in.DecisionStrategy != nil {
		in, out := &in.DecisionStrategy, &out.DecisionStrategy
		*out = new(string)
		**out = **in
	}
	if in.KeepDefaults != nil {
		in, out := &in.KeepDefaults, &out.KeepDefaults
		*out = new(bool)
		**out = **in
	}
	if in.PolicyEnforcementMode != nil {
		in, out := &in.PolicyEnforcementMode, &out.PolicyEnforcementMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationParameters.
func (in *AuthorizationParameters) DeepCopy() *AuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(AuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Client) DeepCopyInto(out *Client) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Client.
func (in *Client) DeepCopy() *Client {
	if in == nil {
		return nil
	}
	out := new(Client)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Client) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientList) DeepCopyInto(out *ClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Client, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientList.
func (in *ClientList) DeepCopy() *ClientList {
	if in == nil {
		return nil
	}
	out := new(ClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientObservation) DeepCopyInto(out *ClientObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ResourceServerID != nil {
		in, out := &in.ResourceServerID, &out.ResourceServerID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountUserID != nil {
		in, out := &in.ServiceAccountUserID, &out.ServiceAccountUserID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientObservation.
func (in *ClientObservation) DeepCopy() *ClientObservation {
	if in == nil {
		return nil
	}
	out := new(ClientObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientParameters) DeepCopyInto(out *ClientParameters) {
	*out = *in
	if in.AccessTokenLifespan != nil {
		in, out := &in.AccessTokenLifespan, &out.AccessTokenLifespan
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.AdminURL != nil {
		in, out := &in.AdminURL, &out.AdminURL
		*out = new(string)
		**out = **in
	}
	if in.AuthenticationFlowBindingOverrides != nil {
		in, out := &in.AuthenticationFlowBindingOverrides, &out.AuthenticationFlowBindingOverrides
		*out = make([]AuthenticationFlowBindingOverridesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = make([]AuthorizationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackchannelLogoutRevokeOfflineSessions != nil {
		in, out := &in.BackchannelLogoutRevokeOfflineSessions, &out.BackchannelLogoutRevokeOfflineSessions
		*out = new(bool)
		**out = **in
	}
	if in.BackchannelLogoutSessionRequired != nil {
		in, out := &in.BackchannelLogoutSessionRequired, &out.BackchannelLogoutSessionRequired
		*out = new(bool)
		**out = **in
	}
	if in.BackchannelLogoutURL != nil {
		in, out := &in.BackchannelLogoutURL, &out.BackchannelLogoutURL
		*out = new(string)
		**out = **in
	}
	if in.BaseURL != nil {
		in, out := &in.BaseURL, &out.BaseURL
		*out = new(string)
		**out = **in
	}
	if in.ClientAuthenticatorType != nil {
		in, out := &in.ClientAuthenticatorType, &out.ClientAuthenticatorType
		*out = new(string)
		**out = **in
	}
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClientOfflineSessionIdleTimeout != nil {
		in, out := &in.ClientOfflineSessionIdleTimeout, &out.ClientOfflineSessionIdleTimeout
		*out = new(string)
		**out = **in
	}
	if in.ClientOfflineSessionMaxLifespan != nil {
		in, out := &in.ClientOfflineSessionMaxLifespan, &out.ClientOfflineSessionMaxLifespan
		*out = new(string)
		**out = **in
	}
	if in.ClientSecretSecretRef != nil {
		in, out := &in.ClientSecretSecretRef, &out.ClientSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ClientSessionIdleTimeout != nil {
		in, out := &in.ClientSessionIdleTimeout, &out.ClientSessionIdleTimeout
		*out = new(string)
		**out = **in
	}
	if in.ClientSessionMaxLifespan != nil {
		in, out := &in.ClientSessionMaxLifespan, &out.ClientSessionMaxLifespan
		*out = new(string)
		**out = **in
	}
	if in.ConsentRequired != nil {
		in, out := &in.ConsentRequired, &out.ConsentRequired
		*out = new(bool)
		**out = **in
	}
	if in.ConsentScreenText != nil {
		in, out := &in.ConsentScreenText, &out.ConsentScreenText
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DirectAccessGrantsEnabled != nil {
		in, out := &in.DirectAccessGrantsEnabled, &out.DirectAccessGrantsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayOnConsentScreen != nil {
		in, out := &in.DisplayOnConsentScreen, &out.DisplayOnConsentScreen
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExcludeSessionStateFromAuthResponse != nil {
		in, out := &in.ExcludeSessionStateFromAuthResponse, &out.ExcludeSessionStateFromAuthResponse
		*out = new(bool)
		**out = **in
	}
	if in.ExtraConfig != nil {
		in, out := &in.ExtraConfig, &out.ExtraConfig
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.FrontchannelLogoutEnabled != nil {
		in, out := &in.FrontchannelLogoutEnabled, &out.FrontchannelLogoutEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FrontchannelLogoutURL != nil {
		in, out := &in.FrontchannelLogoutURL, &out.FrontchannelLogoutURL
		*out = new(string)
		**out = **in
	}
	if in.FullScopeAllowed != nil {
		in, out := &in.FullScopeAllowed, &out.FullScopeAllowed
		*out = new(bool)
		**out = **in
	}
	if in.ImplicitFlowEnabled != nil {
		in, out := &in.ImplicitFlowEnabled, &out.ImplicitFlowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoginTheme != nil {
		in, out := &in.LoginTheme, &out.LoginTheme
		*out = new(string)
		**out = **in
	}
	if in.Oauth2DeviceAuthorizationGrantEnabled != nil {
		in, out := &in.Oauth2DeviceAuthorizationGrantEnabled, &out.Oauth2DeviceAuthorizationGrantEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Oauth2DeviceCodeLifespan != nil {
		in, out := &in.Oauth2DeviceCodeLifespan, &out.Oauth2DeviceCodeLifespan
		*out = new(string)
		**out = **in
	}
	if in.Oauth2DevicePollingInterval != nil {
		in, out := &in.Oauth2DevicePollingInterval, &out.Oauth2DevicePollingInterval
		*out = new(string)
		**out = **in
	}
	if in.PkceCodeChallengeMethod != nil {
		in, out := &in.PkceCodeChallengeMethod, &out.PkceCodeChallengeMethod
		*out = new(string)
		**out = **in
	}
	if in.RealmID != nil {
		in, out := &in.RealmID, &out.RealmID
		*out = new(string)
		**out = **in
	}
	if in.RootURL != nil {
		in, out := &in.RootURL, &out.RootURL
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountsEnabled != nil {
		in, out := &in.ServiceAccountsEnabled, &out.ServiceAccountsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StandardFlowEnabled != nil {
		in, out := &in.StandardFlowEnabled, &out.StandardFlowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.UseRefreshTokens != nil {
		in, out := &in.UseRefreshTokens, &out.UseRefreshTokens
		*out = new(bool)
		**out = **in
	}
	if in.UseRefreshTokensClientCredentials != nil {
		in, out := &in.UseRefreshTokensClientCredentials, &out.UseRefreshTokensClientCredentials
		*out = new(bool)
		**out = **in
	}
	if in.ValidRedirectUris != nil {
		in, out := &in.ValidRedirectUris, &out.ValidRedirectUris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WebOrigins != nil {
		in, out := &in.WebOrigins, &out.WebOrigins
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientParameters.
func (in *ClientParameters) DeepCopy() *ClientParameters {
	if in == nil {
		return nil
	}
	out := new(ClientParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSpec) DeepCopyInto(out *ClientSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSpec.
func (in *ClientSpec) DeepCopy() *ClientSpec {
	if in == nil {
		return nil
	}
	out := new(ClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientStatus) DeepCopyInto(out *ClientStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientStatus.
func (in *ClientStatus) DeepCopy() *ClientStatus {
	if in == nil {
		return nil
	}
	out := new(ClientStatus)
	in.DeepCopyInto(out)
	return out
}
