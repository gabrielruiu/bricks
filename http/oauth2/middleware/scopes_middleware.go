// Copyright © 2019 by PACE Telematics GmbH. All rights reserved.
// Created at 2019/03/28 by Florian Hübsch

package middleware

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/pace/bricks/http/oauth2"
)

// RequiredScopes defines the scope each endpoint requires
type RequiredScopes map[string]oauth2.Scope

// ScopesMiddleware contains required scopes for each endpoint
type ScopesMiddleware struct {
	RequiredScopes RequiredScopes
}

// NewScopesMiddleware return a new scopes middleware
func NewScopesMiddleware(scopes RequiredScopes) *ScopesMiddleware {
	return &ScopesMiddleware{RequiredScopes: scopes}
}

// Handler checks if the token extracted from the request's context has the required scope
// for the requested route and returns a 401 response if not.
func (m *ScopesMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routeName := mux.CurrentRoute(r).GetName()
		if oauth2.HasScope(r.Context(), m.RequiredScopes[routeName]) {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, fmt.Sprintf("Forbidden - requires scope %q", m.RequiredScopes[routeName]), http.StatusForbidden)
	})
}
