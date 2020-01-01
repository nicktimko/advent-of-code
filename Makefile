#### 'Opinonated approach to make files'
SHELL := bash
.ONESHELL:
.SHELLFLAGS := -o errexit -o nounset -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >
####

go_sources := $(shell find . -type f -name '*.go')

aoc2019: ${go_sources}
> go build -o aoc2019

clean:
> rm -rf aoc2019

# "make 1" -> solve for day 1
%:
> make aoc2019
> ./aoc2019 -day=$@

test: ${go_sources}
> go test ./...
