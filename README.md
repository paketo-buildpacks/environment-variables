# `gcr.io/paketo-buildpacks/environment-variables`

The Paketo Environment Variables Buildpack is a Cloud Native Buildpack that embeds environment variables into an image.

## Behavior

This buildpack will participate all the following conditions are met

* Any environment variable matching `BPE_*` is set

The buildpack will do the following:

* Modify the launch environment using the `BPE_*` envrionment variables, as described in the Configuration section below

## Configuration

| Environment Variable   | Description                                                |
| ---------------------- | ---------------------------------------------------------- |
| `BPE_<NAME>`          | set `$NAME` to value (same as override)                    |
| `BPE_APPEND_<NAME>`   | append value to `$NAME`                                    |
| `BPE_DEFAULT_<NAME>`  | set default value for `$NAME`                              |
| `BPE_DELIM_<NAME>`    | set delimeter to use when appending or prepending to $NAME |
| `BPE_OVERRIDE_<NAME>` | set `$NAME` to value                                       |
| `BPE_PREPEND_<NAME>`  | prepend value to `$NAME`                                   |

The default delimiter is an empty string, i.e. there is no default delimiter.

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
