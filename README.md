# robots-disallow
robots-disallow implements a http server serving content of a robots.txt file disallowing all robots.

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

