package lambda_proc

import (
	"encoding/json"
)

// AWS API Gateway Proxy
type APIGatewayProxyEvent struct {
	Body                  *json.RawMessage  `json:"body"`
	Resource              string            `json:"resource"`
	RequestContext        RequestContext    `json:"requestContext"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	Headers               map[string]string `json:"headers"`
	PathParameters        map[string]string `json:"pathParameters"`
	HTTPMethod            string            `json:"httpMethod"`
	StageVariables        map[string]string `json:"stageVariables"`
	Path                  string            `json:"path"`
}

type RequestContext struct {
	ResourceID   string          `json:"resourceId"`
	APIid        string          `json:"apiId"`
	ResourcePath string          `json:"resourcePath"`
	HTTPMethod   string          `json:"httpMethod"`
	RequestID    string          `json:"requestId"`
	AccountID    string          `json:"accountId"`
	Identity     RequestIdentity `json:"identity"`
	Stage        string          `json:"stage"`
}

type RequestIdentity struct {
	APIKey                        *string `json:"apiKey"`
	UserARN                       *string `json:"userArn"`
	CognitoAuthenticationType     *string `json:"cognitoAuthenticationType"`
	Caller                        *string `json:"caller"`
	UserAgent                     *string `json:"userAgent"`
	User                          *string `json:"user"`
	CognitoIdentityPoolID         *string `json:"cognitoIdentityPoolId"`
	CognitoIdentityID             *string `json:"cognitoIdentityId"`
	CognitoAuthenticationProvider *string `json:"cognitoAuthenticationProvider"`
	SourceIP                      *string `json:"sourceIp"`
	AccountID                     *string `json:"accountId"`
}

type APIGatewayProxyResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}
