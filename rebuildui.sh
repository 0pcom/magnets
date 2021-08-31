#!/usr/bin/bash
#recompile hugo sources
hugo
#rebuild static resources
statik -f -src=./public
go generate
