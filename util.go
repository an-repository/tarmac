/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "net/http"

func applyMiddleware(handler HandlerFunc, middlewares ...MiddlewareFunc) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func getPath(r *http.Request) string {
	if path := r.URL.RawPath; path != "" {
		return path
	}

	return r.URL.Path
}

/*
######################################################################################################## @(^_^)@ #######
*/
