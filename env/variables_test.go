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

package env_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/environment-variables/env"
)

func testVariables(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect    = NewWithT(t).Expect
		ctx       libcnb.BuildContext
		variables *env.Variables
	)

	it.Before(func() {
		var err error
		ctx.Layers.Path, err = ioutil.TempDir("", "application-layers")
		Expect(err).NotTo(HaveOccurred())
		envVars := map[string]string{
			"some-key":  "some-val",
			"other-key": "other-val",
		}
		variables = env.NewVariables(envVars)
	})

	it.After(func() {
		Expect(os.RemoveAll(ctx.Layers.Path)).To(Succeed())
	})

	it("contributes to launch environment", func() {
		layer, err := ctx.Layers.Layer("test-layer")
		Expect(err).NotTo(HaveOccurred())

		layer, err = variables.Contribute(layer)
		Expect(err).NotTo(HaveOccurred())
		Expect(layer.Launch).To(BeTrue())
		Expect(layer.LaunchEnvironment).To(Equal(variables.Environment))
		Expect(layer.Metadata).To(HaveKey("variables"))
		Expect(layer.Metadata["variables"]).To(HaveKeyWithValue("some-key", "some-val"))
		Expect(layer.Metadata["variables"]).To(HaveKeyWithValue("other-key", "other-val"))
	})
}
