package model

// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
type OpenIDProviderMetaData struct {
	// Required
	Issuer string `json:"issuer"`
	// Required
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// Required
	TokenEndpoint string `json:"token_endpoint"`
	// Recommended
	UserinfoEndpoint string `json:"userinfo_endpoint"`
	// Recommended
	ScopesSupported []string `json:"scopes_supported"`
	// Required
	JwksURI string `json:"jwks_uri"`
	// Required
	ResponseTypesSupported []string `json:"response_types_supported"`
	// Required
	SubjectTypesSupported []string `json:"subject_types_supported"`
	// Required
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	// Recommended
	ClaimsSupported []string `json:"claims_supported"`
	// Recommended
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}
