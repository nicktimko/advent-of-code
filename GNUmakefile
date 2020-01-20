#### 'Opinonated approach to make files' https://tech.davis-hansson.com/p/make/
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
MAKEFLAGS += --no-print-directory

go_sources := $(shell find . -type f -name '*.go')

aoc2019: ${go_sources}
> go build -o aoc2019

outputs/%.txt: aoc2019
> mkdir -p outputs
> filename=$$(basename -- $@)
> daynum="$${filename%.*}"
> ./$< -day=$${daynum} > $@

##############################
# phony targets
.PHONY: all clean test outputs

all: clean test outputs

clean:
> rm -rf aoc2019
> rm -rf outputs/

outputs: aoc2019
> make outputs/7.txt &
> make outputs/6.txt &
> make outputs/5.txt &
> make outputs/4.txt &
> make outputs/3.txt &
> make outputs/2.txt &
> make outputs/1.txt

test: ${go_sources}
> go test ./...

print%: aoc2019
> target=$@
> daynum=$${target:5}
> ./$< -day=$${daynum}
