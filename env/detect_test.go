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
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/environment-variables/env"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.DetectContext
		detect env.Detect
	)

	context("$BPE_*", func() {
		it("detects", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{Pass: false}))
		})
	})

	context("without $BPE_*", func() {
		it.Before(func() {
			ctx.Platform.Environment = map[string]string{"BPE_SOME_KEY": "some-val"}
		})

		it("fails to detect", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{
				Pass: true,
				Plans: []libcnb.BuildPlan{
					{
						Provides: []libcnb.BuildPlanProvide{
							{Name: "environment-variables"},
						},
						Requires: []libcnb.BuildPlanRequire{
							{Name: "environment-variables"},
						},
					},
				},
			}))
		})
	})
}
