/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "net/http"

type Response struct {
	http.ResponseWriter
	status    int
	committed bool
}

func newResponse() *Response {
	return &Response{}
}

func (r *Response) reset(w http.ResponseWriter) {
	r.ResponseWriter = w
	r.status = 0
	r.committed = false
}

func (r *Response) WriteHeader(status int) {
	if !r.committed {
		r.ResponseWriter.WriteHeader(status)
		r.status = status
		r.committed = true
	}
}

/*
######################################################################################################## @(^_^)@ #######
*/
