# gURL
Essentially a shitty cURL CLI implementation in Go which basically simulates HTTP requests. I wrote this in order to familiarize myself with Go, as well as to refresh myself on networking concepts.

## Installation
```
go get -u github.com/bjma/gurl
```

Clone this repository:
```
git clone https://github.com/bjma/gurl.git
cd gurl
go install
```

## Usage
Run the Makefile using `make` to compile the executable. Then, simply run `gurl` like the following:
```
./gurl [URL] [OPTIONS -f]
```

To run the project globally, you can create an alias in your `.bash_profile`:
```bash
# ~/.bash_profile
alias gurl="/path/to/project/gurl "
```

### Options
* `url` - Specify URL (defaulted to `argv[1]`)
* `X` - Specify METHOD (defaulted to `GET`)
* `d` - Read data from a `string` or file prefixed by `@` **TODO**
* `o` - Write output to file (defaulted to `stdout`)
* `H` - Set header
* `s` - Suppress header output
* `v` - Debugging
