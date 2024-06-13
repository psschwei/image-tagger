# image-tagger
Get a list of tags from an OCI registry

## Usage

```bash
export REGISTRY=https://$(oc registry info)
export TOKEN=$(oc whoami -t)
export PROJECT=<your project/image> # will check as prefix

go run .
```
