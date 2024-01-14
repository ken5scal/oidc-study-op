package model

// OAuth2AuthzRequest represents an OAuth2 authorization request.
// https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthorizationEndpoint
type OAuth2AuthzRequest struct {
	// required
	Scope         []string
	response_type string
	client_id     string
	redirect_uri  string

	// optional
	state string
}

type OIDCAuthzRequest struct {
	OAuth2AuthzRequest

	// optional
	response_mode string
	nonce         string
	display       string // formatは決まっている
	prompt        string // formatは決まっている
	max_age       string
	ui_locales    string // 決まっている
	id_token_hint string
	login_hint    string
	acr_values    string
}
