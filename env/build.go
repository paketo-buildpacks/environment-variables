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
	"fmt"
	"strings"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak/bard"
)

const (
	// Prefix is used to select environment variables that should be contributed to the image
	Prefix             = "BPE_"
	PrefixTypeAppend   = Prefix + "APPEND_"
	PrefixTypeDefault  = Prefix + "DEFAULT_"
	PrefixTypeDelim    = Prefix + "DELIM_"
	PrefixTypeOverride = Prefix + "OVERRIDE_"
	PrefixTypePrepend  = Prefix + "PREPEND_"
)

type Build struct {
	Logger bard.Logger
}

type modType struct {
	prefix string
	suffix string
}

// modTypes maps environment variable prefixes to environment file suffixes
var modTypes = []modType{
	{prefix: PrefixTypeAppend, suffix: ".append"},
	{prefix: PrefixTypeDefault, suffix: ".default"},
	{prefix: PrefixTypeDelim, suffix: ".delim"},
	{prefix: PrefixTypeOverride, suffix: ".override"},
	{prefix: PrefixTypePrepend, suffix: ".prepend"},
	{prefix: Prefix, suffix: ""}, // untyped prefix goes last because it is a substring of the typed prefixes
}

func (b Build) Build(context libcnb.BuildContext) (libcnb.BuildResult, error) {
	b.Logger.Title(context.Buildpack)
	b.LogConfiguration()

	envVars := map[string]string{}
	for k, v := range context.Platform.Environment {
		v := v
		for _, modType := range modTypes {
			if strings.HasPrefix(k, modType.prefix) {
				envVars[strings.TrimPrefix(k, modType.prefix)+modType.suffix] = v
				break
			}
		}
	}
	variables := NewVariables(envVars)
	variables.Logger = b.Logger
	return libcnb.BuildResult{
		Layers: []libcnb.LayerContributor{variables},
	}, nil
}

func (b Build) LogConfiguration() {
	var nameLength int
	for _, mt := range modTypes {
		if l := len(key(mt.prefix)); l > nameLength {
			nameLength = l
		}
	}
	b.Logger.Header(fmt.Sprint("Launch Configuration:"))
	b.Logger.Bodyf("$%s\t prepend value to $NAME, delimiting with OS path list separator", pad(key(Prefix), nameLength))
	b.Logger.Bodyf("$%s\t append value to $NAME", pad(key(PrefixTypeAppend), nameLength))
	b.Logger.Bodyf("$%s\t set default value for $NAME", pad(key(PrefixTypeDefault), nameLength))
	b.Logger.Bodyf("$%s\t set delimeter to use when appending or prepending to $NAME", pad(key(PrefixTypeDelim), nameLength))
	b.Logger.Bodyf("$%s\t set $NAME to value", pad(key(PrefixTypeOverride), nameLength))
	b.Logger.Bodyf("$%s\t prepend value to $NAME", pad(key(PrefixTypePrepend), nameLength))
}

func key(prefix string) string {
	return prefix + `<NAME>`
}

func pad(s string, length int) string {
	for {
		if len(s) == length {
			return s
		}
		s = s + " "
	}
}
