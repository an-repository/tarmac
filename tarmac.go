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
	Tarmac struct {
		router *Router
	}

	HandlerFunc    func(*Context) error
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

var _allMethods = []string{
	http.MethodConnect,
	http.MethodDelete,
	http.MethodGet,
	http.MethodHead,
	http.MethodOptions,
	http.MethodPatch,
	http.MethodPost,
	http.MethodPut,
	http.MethodTrace,
}

func New() *Tarmac {
	return &Tarmac{
		router: NewRouter(),
	}
}

func (t *Tarmac) Router() *Router {
	return t.router
}

func (t *Tarmac) Add(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.router.Add(
		method,
		path,
		func(c *Context) error {
			return applyMiddleware(handler, middlewares...)(c)
		},
	)
}

func (t *Tarmac) Delete(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.Add(http.MethodDelete, path, handler, middlewares...)
}

func (t *Tarmac) Get(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.Add(http.MethodGet, path, handler, middlewares...)
}

func (t *Tarmac) Post(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.Add(http.MethodPost, path, handler, middlewares...)
}

func (t *Tarmac) Put(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.Add(http.MethodPut, path, handler, middlewares...)
}

func (t *Tarmac) Any(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	for _, m := range _allMethods {
		if err := t.Add(m, path, handler, middlewares...); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tarmac) Group(prefix string, middlewares ...MiddlewareFunc) *Group {
	g := &Group{
		prefix: prefix,
		tarmac: t,
	}

	g.Use(middlewares...)

	return g
}

/*
######################################################################################################## @(^_^)@ #######
*/
