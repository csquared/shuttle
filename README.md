# `shuttle`

Shuttle newline delimited data from STDIN to http(s)

[![Gobuild Download](http://gobuild.io/badge/github.com/csquared/shuttle/downloads.svg)](http://gobuild.io/github.com/csquared/shuttle)


## usage

    echo foo | shuttle http://localhost:4567

Sends `foo\n` with a mime-type of `text/plain` to http://localhost:4567
