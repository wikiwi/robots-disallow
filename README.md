# robots-disallow
robots-disallow implements a http server serving content of a robots.txt file disallowing all robots.

## Usage
    robots-disallow serve [flags]

    Flags:
          --listen string         address to listen on [$ROBOTS_DISALLOW_LISTEN] (default "0.0.0.0:8080")
## Example
    docker run -p 8080:8080 wikiwi/robots-disallow serve

## Output
    curl localhost:8080
    User-agent: *
    Disallow: /

## Docker Hub
Automated build is available at the [Docker Hub](https://hub.docker.com/r/wikiwi/robots-disallow).

