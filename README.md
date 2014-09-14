# `shuttle`

Shuttle newline delimited data from STDIN to http(s)


## usage

    echo foo | shuttle http://localhost:4567

Sends `foo\n` with a mime-type of `text/plain` to http://localhost:4567
