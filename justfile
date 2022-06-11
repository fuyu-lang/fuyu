# When set to true all commands are run with verbose output.
verbose := 'false'

# Sets the flag to use enale verbose output in go commands.
go_verbose := if verbose == 'true' { '-x' } else { '' }

# The directory where all build outputs are written.
output_dir := './output/'

# list all recipes
@default:
  just --list

# format all code
@fmt:
  go fmt {{go_verbose}} ./...

# generate all code
@generate:
  go generate {{go_verbose}} ./...

# build all projects
@build:
  mkdir {{output_dir}}
  go build -o {{output_dir}} {{go_verbose}} ./...

# run all tests
@test:
  go test {{go_verbose}} ./...

# clean all build products
@clean:
  rm -rf {{output_dir}}
