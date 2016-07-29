# robots-disallow
robots-disallow implements a http server serving content of a robots.txt file disallowing all robots.

[![Build Status Widget]][Build Status] [![Coverage Status Widget]][Coverage Status] [![Code Climate Widget]][Code Climate] [![MicroBadger Image Widget]][MicroBadger URL]

[Build Status]: https://travis-ci.org/wikiwi/robots-disallow
[Build Status Widget]: https://travis-ci.org/wikiwi/robots-disallow.svg?branch=master
[Coverage Status]: https://coveralls.io/github/wikiwi/robots-disallow?branch=master
[Coverage Status Widget]: https://coveralls.io/repos/github/wikiwi/robots-disallow/badge.svg?branch=master
[Code Climate]: https://codeclimate.com/github/wikiwi/robots-disallow
[Code Climate Widget]: https://codeclimate.com/github/wikiwi/robots-disallow/badges/gpa.svg
[MicroBadger URL]: http://microbadger.com/#/images/wikiwi/robots-disallow
[MicroBadger Image Widget]: https://images.microbadger.com/badges/image/wikiwi/robots-disallow.svg


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

