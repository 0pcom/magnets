#!/usr/bin/bash
#recompile hugo sources
hugo --config="config.toml" --baseURL="https://magnetosphereelectronicsurplus.com" -d public1
#rebuild static resources
statik -f -src=./public1
go generate
hugo --config="config.toml" --baseURL="https://magnetosphere.net" -d public
statik -f -src=./public
go generate
