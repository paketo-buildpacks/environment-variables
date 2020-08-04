/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package env

import (
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

type Variables struct {
	libcnb.Environment
	LayerContributor libpak.LayerContributor
	Logger           bard.Logger
}

func NewVariables(vars map[string]string) *Variables {
	return &Variables{
		Environment: vars,
		LayerContributor: libpak.NewLayerContributor("Environment Variables", map[string]interface{}{
			"variables": vars,
		}),
	}
}

func (v Variables) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	v.LayerContributor.Logger = v.Logger

	return v.LayerContributor.Contribute(layer, func() (libcnb.Layer, error) {
		layer.Launch = true
		layer.LaunchEnvironment = v.Environment
		return layer, nil
	})
}

func (Variables) Name() string {
	return "environment-variables"
}
