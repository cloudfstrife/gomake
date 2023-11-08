# gomake 

makefile for go language 

## Features

* single and multi executable program project 
* Multi-platform bulid
* code format
* SonarQube integration
* benchmark test

## Status

The project is in v1.0 version.

## Getting started

* create project 
* copy makefile.single or makefile.multi to project folder as makefile
* modify makefile
* do something by run make command

### command

* `make build` build current platform executable file into $DIST_FOLDER  (default)
* `make format` format go source file 
* `make test` run unit test
* `make benchmark` run benchmark test
* `make sonar` SonarQube integration
* `make all` build all platform executable file
* `make clean` remove $DIST_FOLDER and SonarQube report file 
* `make container` build docker image
* `make push-container` push docker image to docker repository

## Documentation

## Contributing

If you are interested in contributing, please checkout [CONTRIBUTING.md](./CONTRIBUTING.md).
We welcome any code or non-code contribution!

## Licensing

licensed by the MIT License. See [LICENSE](./LICENSE) for the full license text.
