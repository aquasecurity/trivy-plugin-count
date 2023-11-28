# trivy-output-plugin-count
Example of Trivy output plugin

## Installation
```shell
trivy plugin install github.com/aquasecurity/trivy-output-plugin-count
```

## Usage

```shell
trivy <target> --format json --output plugin=count [--output-plugin-arg plugin_flags] <target_name>
```

OR

```shell
trivy <target> -f json <target_name> | trivy count [plugin_flags]
```

## Examples

```shell
trivy image -f json -o plugin=count --output-plugin-arg "--published-after=2023-11-01" debian:12
```

is equivalent to:

```shell
trivy image -f json debian:12 | trivy count --published-after=2023-11-01
```
