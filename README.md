# Usable type definitions for Go language

Introduction
------------

Some of Go's built-in types are lacking required functionality. For example,
`time.Duration` value cannot be easily marshaled / unmarshaled to / from
JSON ot YAML document.  
This package provides few convenient type helpers to address the issue.

Installation and usage
----------------------

The import path for the package is *github.com/ipitsyn/gotypes*.

To install it, run:

    go get github.com/ipitsyn/gotypes
