/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "net/http"

type Group struct {
	tarmac      *Tarmac
	prefix      string
	middlewares []MiddlewareFunc
}

func newGroup(t *Tarmac, prefix string) *Group {
	return &Group{
		tarmac: t,
		prefix: prefix,
	}
}

func (g *Group) Use(middlewares ...MiddlewareFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
	if len(g.middlewares) == 0 {
		return
	}

	/* FIXME ?
	// Allow all requests to reach the group as they might get dropped if router
	// doesn't find a match, making none of the group middleware process.
	g.Any("", NotFoundHandler)
	g.Any("/*", NotFoundHandler)
	*/
}

func (g *Group) Add(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	m := make([]MiddlewareFunc, 0, len(g.middlewares)+len(middlewares))
	m = append(m, g.middlewares...)
	m = append(m, middlewares...)

	return g.tarmac.Add(method, g.prefix+path, handler, m...)
}

func (g *Group) Delete(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return g.Add(http.MethodDelete, path, handler, middlewares...)
}

func (g *Group) Get(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return g.Add(http.MethodGet, path, handler, middlewares...)
}

func (g *Group) Post(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return g.Add(http.MethodPost, path, handler, middlewares...)
}

func (g *Group) Put(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return g.Add(http.MethodPut, path, handler, middlewares...)
}

func (g *Group) Any(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	for _, m := range _allMethods {
		if err := g.Add(m, path, handler, middlewares...); err != nil {
			return err
		}
	}

	return nil
}

func (g *Group) Group(prefix string, middlewares ...MiddlewareFunc) *Group {
	m := make([]MiddlewareFunc, 0, len(g.middlewares)+len(middlewares))
	m = append(m, g.middlewares...)
	m = append(m, middlewares...)

	return g.tarmac.Group(g.prefix+prefix, m...)
}

/*
######################################################################################################## @(^_^)@ #######
*/
