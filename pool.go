/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"sync"

	"github.com/an-repository/tracer"
)

type pool struct {
	sp *sync.Pool
}

func newPool() *pool {
	return &pool{
		sp: &sync.Pool{
			New: func() any {
				tracer.Send("[tarmac] new context instance") //.........................................................
				return &Context{}
			},
		},
	}
}

func (p *pool) get() *Context {
	c := p.sp.Get().(*Context)
	c.reset()

	return c
}

func (p *pool) put(c *Context) {
	p.sp.Put(c)
}

/*
######################################################################################################## @(^_^)@ #######
*/
