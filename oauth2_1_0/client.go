// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package oauth2_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetUserTokenRequest struct {
	// 应用id
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// 应用密码
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
	// OAuth 2.0 临时授权码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// OAuth 2.0 刷新令牌
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// 分为authorization_code和refresh_token。使用授权码换token，传authorization_code；使用刷新token换用户token，传refresh_token
	GrantType *string `json:"grantType,omitempty" xml:"grantType,omitempty"`
}

func (s GetUserTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserTokenRequest) SetClientId(v string) *GetUserTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetUserTokenRequest) SetClientSecret(v string) *GetUserTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *GetUserTokenRequest) SetCode(v string) *GetUserTokenRequest {
	s.Code = &v
	return s
}

func (s *GetUserTokenRequest) SetRefreshToken(v string) *GetUserTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *GetUserTokenRequest) SetGrantType(v string) *GetUserTokenRequest {
	s.GrantType = &v
	return s
}

type GetUserTokenResponseBody struct {
	// accessToken
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// refreshToken
	RefreshToken *string `json:"refreshToken,omitempty" xml:"refreshToken,omitempty"`
	// 超时时间
	ExpireIn *int64 `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

func (s GetUserTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBody) SetAccessToken(v string) *GetUserTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetUserTokenResponseBody) SetRefreshToken(v string) *GetUserTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *GetUserTokenResponseBody) SetExpireIn(v int64) *GetUserTokenResponseBody {
	s.ExpireIn = &v
	return s
}

type GetUserTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponse) SetHeaders(v map[string]*string) *GetUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserTokenResponse) SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) GetUserToken(request *GetUserTokenRequest) (_result *GetUserTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserTokenResponse{}
	_body, _err := client.GetUserTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserTokenWithOptions(request *GetUserTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		body["clientSecret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		body["refreshToken"] = request.RefreshToken
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		body["grantType"] = request.GrantType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetUserTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserToken"), tea.String("oauth2_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/oauth2/userAccessToken"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
