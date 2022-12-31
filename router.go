/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "errors"

type router struct {
}

func newRouter() *router {
	return &router{}
}

func (r *router) add(method, path string, handler HandlerFunc) error {
	if handler == nil {
		return errors.New("handler cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	return nil
}

func (r *router) find(c *Context, method, path string) HandlerFunc {
	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
