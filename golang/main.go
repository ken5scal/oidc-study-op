package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ken5scal/oidc-study-op/model"
)

func main() {
	http.HandleFunc("/.well-known/openid-configuration", discoveryHandler)
	http.ListenAndServe(":8080", nil)
}

func discoveryHandler(w http.ResponseWriter, r *http.Request) {
	baseUrl := "https://" + r.Host
	response := model.OpenIDProviderMetaData{
		Issuer:                 baseUrl,
		AuthorizationEndpoint:  baseUrl + "/oauth2/authorize",
		TokenEndpoint:          baseUrl + "/oauth2/token",
		JwksURI:                baseUrl + "/.well-known/jwks.json",
		ResponseTypesSupported: []string{"code", "id_token", "token id_token"},
		SubjectTypesSupported:  []string{"public"},

		IDTokenSigningAlgValuesSupported: []string{"RS256"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
