/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import "errors"

type (
	Route struct {
		Name   string
		Method string
		Path   string
	}

	Router struct {
	}
)

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method, path string, handler HandlerFunc) error {
	if handler == nil {
		return errors.New("handler cannot be nil") /////////////////////////////////////////////////////////////////////
	}

	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
