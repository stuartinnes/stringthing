# stringthing - a handy CLI string toolkit 

A smashing little application for string generation and validation.

---
## Motivation

<span style="color:#83c462">
> Be you <br />
> Be in a hurry <br />
> Realise you can't learn grep in 3 minutes <br />
> PaNiK.mp3 <br />
> Use <b>stringthing</b> instead <br />
> Kalm.wav
</span>

Everyone can be a command line ninja but sometimes getting stuff done without dusting off your **grep**, **awk** and **sed** skills is more important.

### Usage 
Currently stringthing supports **format checking**, **random string generation** and **v4 UUID generation**. 

```shell
> stringthing -h 

NAME:
   stringthing - A handy CLI string toolkit 

USAGE:
   stringthing [global options] command [command options] [arguments...]

COMMANDS:
   is       format checks
   random   generate random strings
   uuid     generate V4 uuids
   help, h  Shows a list of commands or help for one command 
```

### Example format check (JSON)

```shell
# check if a value is valid JSON 
stringthing is json '{"isThis":"JSON?"}' # exit code 0

# check if a value is valid JSON with verbose output 
stringthing is json -v '{"validJSON?"}' # exit code 1 and "false" printed to stdout

# check if a file is valid JSON
cat somefile.json | xargs stringthing is json 

# use this a lot? create an alias 
alias isJSON="stringthing is json"
isJSON '{"isThis":"JSON?"}'
```

A complete list of available string [formats](./formats.md) 

### Example random string 

```shell
# default 
stringthing random # outputs something like "RDuP4+6!^<7km*\3Z81j.75*"

# custom string composition 
stringthing random -c 5 -d 5 -s 0 # outputs something like "789LGesn28"
```

### Example v4 UUIDs

```shell 
#default 
stringthing uuid # outputs something like "fb80b3e9-7b89-4441-ab8c-b4b84cd94b07"

#short 
stringthing uuid -s # outputs something like "YQbsixE5a4rfAFgMT84eaJ"
```

### Verbosity 

By default string thing format checks will fail with a non-zero error code. This is fine in some cases however you may wish to 
get the status printed to stdout. 

```shell
# set and forget environment variable 
export ST_VERBOSE=true

# case by case using the -v flag
stringthing is json -v ...
```



## Installation

**Go** 

```shell 
go install github.com/stuartinnes/stringthing
```

**Homebrew**

```shell 
brew install stuartinnes/tap/stringthing
```

### On the shoulders of giants

This little application would not be possible without the packages that powers it

* [asaskevich govalidator](https://github.com/asaskevich/govalidator) the brains behind the validation 
* [Google UUID](https://github.com/google/uuid) Google's awesome UUID generation package
* [lithammer shortuuid](https://github.com/lithammer/shortuuid) a splendid little (see what I did there?) UUID shortener
* [urfave cli](https://github.com/urfave/cli) bringing it all together as a nice little CLI 

### License 

[MIT](./LICENSE)
