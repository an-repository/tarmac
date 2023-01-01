/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "net/http"

type Context struct {
	Request  *http.Request
	Response *Response
}

func newContext() *Context {
	return &Context{
		Response: newResponse(),
	}
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.Request = r
	c.Response.reset(w)
}

func (c *Context) NoContent(status int) error {
	c.Response.WriteHeader(status)
	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
