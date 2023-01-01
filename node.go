/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"strings"
)

const (
	_paramPrefix = ":"
	_wildcardStr = "..."
)

type node struct {
	param    string
	wildcard bool
	handlers map[string]HandlerFunc
	nodes    map[string]*node
}

func newNode(seg string) *node {
	if seg == _wildcardStr {
		return &node{
			wildcard: true,
		}
	}

	if strings.HasPrefix(seg, _paramPrefix) {
		return &node{
			param: strings.TrimPrefix(seg, _paramPrefix),
		}
	}

	return &node{}
}

func (n *node) next(seg string) *node {
	if n.nodes == nil {
		n.nodes = make(map[string]*node)
		next := newNode(seg)
		n.nodes[seg] = next

		return next
	}

	next, ok := n.nodes[seg]
	if !ok {
		next = newNode(seg)
		n.nodes[seg] = next
	}

	return next
}

func (n *node) add(method string, handler HandlerFunc) {
	if n.handlers == nil {
		n.handlers = make(map[string]HandlerFunc)
	}

	n.handlers[strings.ToUpper(method)] = handler
}

func (n *node) match(c *Context, path string) *node {
	cur := n
	end := false

	var seg string

	for !end {
		i := strings.Index(path, "/")

		if i == -1 {
			seg = path
			path = ""
			end = true
		} else if i == 0 {
			seg = ""
			path = path[1:]
		} else {
			seg = path[:i]
			path = path[i+1:]
		}

		if cur.nodes == nil {
			return nil
		}

		tmp, ok := cur.nodes[seg]
		if ok {
			cur = tmp
		} else {
			for _, v := range cur.nodes {
				if v.param != "" {
					if end {
						if v.nodes == nil {
							// FIXME c.AddParam(v.param, seg)
							return v
						}
					} else if tmp = v.match(c, path); tmp != nil {
						// FIXME c.AddParam(v.param, seg)
						return tmp
					}
				} else if v.wildcard && v.nodes == nil {
					// FIXME c.AddParam(_wildcardStr, strings.TrimRight(seg+"/"+path, "/"))
					return v
				}
			}

			return nil
		}
	}

	return cur
}

func (n *node) allowedMethods() []string {
	var out []string

	if n.handlers != nil {
		out = make([]string, len(n.handlers))

		i := 0
		for m := range n.handlers {
			out[i] = m
			i++
		}
	}

	return out
}

/*
######################################################################################################## @(^_^)@ #######
*/
