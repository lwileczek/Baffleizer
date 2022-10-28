# Bafflizer
The Bafflizer performs stupefaction to a file by increasing the perplexity of the variable and making confusing to read. 
Thus, by obfuscating all of the variables and imports, the code is baffling.

## Details
Currently, this program can be used to discombobulate a python file. 
This is still in active development and is in the alpha stage.
Please see the roadmap or issues to see plans for development.
### Goal
The overall goal of this project is to take a file or project and make it as illegible as possible while retaining functionality.

## Why use it

 - Need to someone code but don't want them to understand it?
 - Afraid someone will look over your shoulder and steal your hard work?
 - Want to scare a friend and make him think his code is bonkers?

## How to use

You can pass a file directly to the baffleizer and it'll spit the output to stdout.
Alternatively, you can use the `--output` to provide an output file.

```bash
$ baffle main.py
```

## Install
Currently, you have to build the binary yourself. You'll need Go 1.16+ and an internet connection.

```bash
$ go build -o baffle main.go
```

If you plan to bewilder a bunch of files you can add the binary to your PATH after it's build. 
On linux or OSX that might look like `sudo mv baffle /usr/local/bin`

## What this is not
A linter or fixer. If you put a broken code file in, this will not fix it or tell you how to improve it. 

> Garbage in, garbage out


## Config
...Under Construction. Show all config options

## Examples
...Under construciton


<!--
    note to self do a release with
    https://goreleaser.com/quick-start/

    Seems lit
-->
