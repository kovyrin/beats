// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package enterprisesearch

// Config defines the structure for the Enterprise Search module configuration options
type Config struct {
	XPackEnabled bool `config:"xpack.enabled"`
}

// DefaultConfig returns the default configuration for the Enterprise Search module
func DefaultConfig() Config {
	return Config{
		XPackEnabled: false,
	}
}
