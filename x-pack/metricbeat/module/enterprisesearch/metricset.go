// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package enterprisesearch

import (
	"github.com/elastic/beats/v7/metricbeat/mb"
)

// MetricSet can be used to build other metricsets within the Enterprise Search module.
type MetricSet struct {
	mb.BaseMetricSet
	XPackEnabled bool
}

// NewMetricSet creates a metricset that can be used to build other metricsets
// within the Enterprise Search module.
func NewMetricSet(base mb.BaseMetricSet) (*MetricSet, error) {
	config := DefaultConfig()
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		base,
		config.XPackEnabled,
	}, nil
}
