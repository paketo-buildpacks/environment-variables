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
	"fmt"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/environment-variables/v4/env"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx      libcnb.BuildContext
		build    env.Build
		modTypes = map[string]string{
			env.Prefix:             "",
			env.PrefixTypeAppend:   ".append",
			env.PrefixTypeDefault:  ".default",
			env.PrefixTypeDelim:    ".delim",
			env.PrefixTypeOverride: ".override",
			env.PrefixTypePrepend:  ".prepend",
		}
		result libcnb.BuildResult
	)

	for prefix, suffix := range modTypes {
		context(fmt.Sprintf("$%s_*", prefix), func() {

			it.Before(func() {
				var err error
				ctx.Platform.Environment = map[string]string{
					prefix + "SOME_KEY": "some-val",
				}
				result, err = build.Build(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result.Layers)).To(Equal(1))
			})

			it(fmt.Sprintf("adds *%s to env", suffix), func() {
				environ := result.Layers[0].(*env.Variables).Environment
				Expect(environ["SOME_KEY"+suffix]).To(Equal("some-val"))
			})
		})
	}
}
