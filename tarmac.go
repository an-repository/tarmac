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
		bMiddlewares []MiddlewareFunc
		aMiddlewares []MiddlewareFunc
		pool         *pool
		router       *router
		server       *server
	}

	HandlerFunc    func(*Context) error
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

var (
	_allMethods = []string{
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

	NotFoundHandler = func(*Context) error {
		return NewStatusError(http.StatusNotFound)
	}
)

func New() *Tarmac {
	return &Tarmac{
		pool:   newPool(),
		router: newRouter(),
	}
}

func (t *Tarmac) UseBefore(middlewares ...MiddlewareFunc) {
	t.bMiddlewares = append(t.bMiddlewares, middlewares...)
}

func (t *Tarmac) UseAfter(middlewares ...MiddlewareFunc) {
	t.aMiddlewares = append(t.aMiddlewares, middlewares...)
}

func (t *Tarmac) Add(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {
	return t.router.add(
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
	g := newGroup(t, prefix)
	g.Use(middlewares...)

	return g
}

func (t *Tarmac) errorHandler(c *Context, err error) {
	if c.Response.committed {
		return
	}

	e, ok := err.(*Error)
	if !ok {
		e = NewError(http.StatusInternalServerError, err)
	}

	if c.Request.Method == http.MethodHead {
		_ = c.NoContent(e.Status)
	}
	/* FIXME
	 else {
		_ = c.JSON(e)
	}
	*/

	if c.Response.committed {
		return
	}

	c.Response.WriteHeader(e.Status)
}

func (t *Tarmac) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := t.pool.get(w, r)
	defer t.pool.put(c)

	var h HandlerFunc

	if t.bMiddlewares == nil {
		h = applyMiddleware(t.router.find(c, r.Method, getPath(r)), t.aMiddlewares...)
	} else {
		h = func(c *Context) error {
			return applyMiddleware(t.router.find(c, r.Method, getPath(r)), t.aMiddlewares...)(c)
		}

		h = applyMiddleware(h, t.bMiddlewares...)
	}

	if err := h(c); err != nil {
		t.errorHandler(c, err)
	}
}

/*
######################################################################################################## @(^_^)@ #######
*/
