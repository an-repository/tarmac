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
	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
