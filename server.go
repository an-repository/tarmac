/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"context"
	"net/http"
)

type server struct {
	cfg    *Config
	server *http.Server
}

func (t *Tarmac) NewServer(cfg *Config) error {
	s := &server{
		cfg: cfg,
		server: &http.Server{
			Addr:              cfg.Addr,
			IdleTimeout:       cfg.IdleTimeout,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			Handler:           t,
			ReadHeaderTimeout: 0,
		},
	}

	tlsConfig, err := cfg.tlsConfig()
	if err != nil {
		return err
	}

	s.server.TLSConfig = tlsConfig

	t.server = s

	return nil
}

func (s *server) Start() error {
	var err error

	if s.cfg.tls {
		err = s.server.ListenAndServeTLS(s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		err = s.server.ListenAndServe()
	}

	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (s *server) Stop(ctx context.Context) error {
	s.server.SetKeepAlivesEnabled(false)
	return s.server.Shutdown(ctx)
}

/*
######################################################################################################## @(^_^)@ #######
*/
