# robots-disallow
robots-disallow implements a http server serving content of a robots.txt file disallowing all robots.

[![Build Status](https://travis-ci.org/wikiwi/robots-disallow.svg?branch=travis)](https://travis-ci.org/wikiwi/robots-disallow) [![Coverage Status](https://coveralls.io/repos/github/wikiwi/robots-disallow/badge.svg?branch=master)](https://coveralls.io/github/wikiwi/robots-disallow?branch=master) [![Code Climate](https://codeclimate.com/github/wikiwi/robots-disallow/badges/gpa.svg)](https://codeclimate.com/github/wikiwi/robots-disallow) [![Docker Hub](https://img.shields.io/docker/pulls/wikiwi/robots-disallow.svg)](https://hub.docker.com/r/wikiwi/robots-disallow)

## Usage
    Usage:
      robots-disallow [OPTIONS]

    Application Options:
          --listen=  address to listen on (default: 0.0.0.0:8080) [$ROBOTS_DISALLOW_LISTEN]
      -v, --version  show version number

    Help Options:
      -h, --help     Show this help message

## Example
    docker run -p 8080:8080 wikiwi/robots-disallow

## Output
    curl localhost:8080
    User-agent: *
    Disallow: /

## Docker Hub
Automated build is available at the [Docker Hub](https://hub.docker.com/r/wikiwi/robots-disallow).

