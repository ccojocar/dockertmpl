# dockertmpl

This is a tool which can be used to generate a Dockerfile based on a go template.

## Installation

```bash
go get github.com/ccojocar/dockertmpl
```

Alternatively you can build the docker image by cloning the repository and executing the following command:

```bash
make image
```

## Generate a Dockerfile from a template

You can define a Dockerfile template using the go template syntax as follows:

```bash
cat >Dockerfile.tmpl <<EOF
FROM {{ .BaseImage }}
ENV BIN={{ .Binary }}
COPY build/*-linux-amd64 /go/bin/$BIN
CMD /go/bin/$BIN
>EOF
```

The values can be now defined in a YAML file:
```bash
cat >values.yaml <<EOF
BaseImage: golang:1.9.4-alpine3.7
Binary: dockertmpl
>EOF
```

The Dockerfile can be generated with the following command:

```bash
dockertmpl -valuesFile values.yaml Dockerfile.tmpl
```

## Development

You can execute the tests and build the tool using the default make target:

```bash
make
```

To build and publish the docker image execute:

```bash
make image
make image-push
```
