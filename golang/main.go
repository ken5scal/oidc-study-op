package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/.well-known/openid-configuration", discoveryHandler)
	http.ListenAndServe(":8080", nil)
}

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

func discoveryHandler(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://" + r.Host
	response := OpenIDProviderMetaData{
		Issuer:                           baseUrl,
		AuthorizationEndpoint:            baseUrl + "/oauth2/authorize",
		TokenEndpoint:                    baseUrl + "/oauth2/token",
		JwksURI:                          baseUrl + "/.well-known/jwks.json",
		ResponseTypesSupported:           []string{"code", "id_token", "token id_token"},
		SubjectTypesSupported:            []string{"public"},
		IDTokenSigningAlgValuesSupported: []string{"RS256"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
