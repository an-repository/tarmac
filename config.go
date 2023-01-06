/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package tarmac

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"time"

	"github.com/an-repository/errors"
)

type (
	Config struct {
		Addr         string        `dm:"addr"`
		IdleTimeout  time.Duration `dm:"idle_timeout"`
		ReadTimeout  time.Duration `dm:"read_timeout"`
		WriteTimeout time.Duration `dm:"write_timeout"`
		CertFile     string        `dm:"cert_file"`
		KeyFile      string        `dm:"key_file"`
		CAFile       string        `dm:"ca_file"`
		//
		ErrorLog *log.Logger
		//
		tls bool
	}

	Option func(*Config)
)

func WithAddr(addr string) Option {
	return func(c *Config) {
		c.Addr = addr
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.IdleTimeout = timeout
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.WriteTimeout = timeout
	}
}

func WithCertKeyFiles(certFile, keyFile string) Option {
	return func(c *Config) {
		c.CertFile = certFile
		c.KeyFile = keyFile
	}
}

func WithCAFile(caFile string) Option {
	return func(c *Config) {
		c.CAFile = caFile
	}
}

func WithErrorLogger(logger *log.Logger) Option {
	return func(c *Config) {
		c.ErrorLog = logger
	}
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}

	for _, option := range opts {
		option(c)
	}

	return c
}

func (c *Config) tlsConfig() (*tls.Config, error) {
	if c.CertFile == "" && c.KeyFile == "" {
		return nil, nil
	}

	c.tls = true

	if ok, err := fileExists(c.CertFile); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("this file doesn't exist", "name", c.CertFile) //////////////////////////////////////////
	}

	if ok, err := fileExists(c.KeyFile); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("this file doesn't exist", "name", c.KeyFile) ///////////////////////////////////////////
	}

	var certPool *x509.CertPool
	authType := tls.NoClientCert

	if c.CAFile != "" {
		authType = tls.RequireAndVerifyClientCert

		buf, err := os.ReadFile(c.CAFile)
		if err != nil {
			return nil, errors.WithMessage(err, "unable to read this file", "file", c.CAFile) //////////////////////////
		}

		certPool = x509.NewCertPool()
		certPool.AppendCertsFromPEM(buf)
	}

	cfg := &tls.Config{
		ClientAuth: authType,
		ClientCAs:  certPool,
		MinVersion: tls.VersionTLS12,
	}

	return cfg, nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
