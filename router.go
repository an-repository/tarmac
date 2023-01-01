/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"errors"
	"net/http"
	"strings"
)

type router struct {
	tree *node
}

func newRouter() *router {
	return &router{}
}

func (r *router) findNode(path string) *node {
	node := r.tree

	for _, seg := range strings.Split(path, "/") {
		node = node.next(seg)
	}

	return node
}

func (r *router) add(method, path string, handler HandlerFunc) error {
	if handler == nil {
		return errors.New("handler cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	if path == "" {
		path = "/"
	} else if path[0] != '/' {
		path = "/" + path
	}

	node := r.findNode(path)
	node.add(method, handler)

	return nil
}

func (r *router) find(c *Context, method, path string) HandlerFunc {
	if node := r.tree.match(c, path); node != nil {
		handler, ok := node.handlers[strings.ToUpper(method)]
		if ok {
			return handler
		}

		if am := node.allowedMethods(); len(am) > 0 {
			amStr := strings.Join(append(am, http.MethodOptions), ", ")

			if method == http.MethodOptions {
				return func(c *Context) error {
					c.Response.Header().Add(HeaderAllow, amStr)
					return c.NoContent(http.StatusNoContent)
				}
			}

			return func(c *Context) error {
				c.Response.Header().Add(HeaderAllow, amStr)
				return NewStatusError(http.StatusMethodNotAllowed)
			}
		}
	}

	return NotFoundHandler
}

/*
######################################################################################################## @(^_^)@ #######
*/
