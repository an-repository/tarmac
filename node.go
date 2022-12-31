/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "strings"

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

/*
######################################################################################################## @(^_^)@ #######
*/
