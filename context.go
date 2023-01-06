/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "net/http"

type (
	param struct {
		name  string
		value string
	}

	Context struct {
		Request  *http.Request
		Response *Response
		params   []param
	}
)

func newContext() *Context {
	return &Context{
		Response: newResponse(),
	}
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.Request = r
	c.Response.reset(w)
	c.params = c.params[:0]
}

func (c *Context) addParam(name, value string) {
	c.params = append(c.params, param{name, value})
}

func (c *Context) Param(name string) string {
	for _, p := range c.params {
		if p.name == name {
			return p.value
		}
	}

	return ""
}

func (c *Context) NoContent(status int) error {
	c.Response.WriteHeader(status)
	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
