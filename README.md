# description

golang program to query same url N times at the same time.
available as docker image: https://hub.docker.com/r/darkwind8/goping

# example to run:

go run . --url="https://httpbin.org/ip" --n=2
```
{
  "origin": "123.123.123.123"
}
%!(EXTRA time.Duration=818.584921ms)
{
  "origin": "123.123.123.123"
}
%!(EXTRA time.Duration=819.336826ms)
total time 819.388399ms
```

# Secret motive
## The program actually served to learn scripting basics for infra scripting purposes

- [x] tried publishing to public docker registry
- [x] learned how to make parallel requests for performance testing
- [x] tried golang testing
- [x] tried how to make simple web server for testing purposes
- [x] tried templating language
- [x] tried publishing golang program to github to be reused as installable dependency to other golang programs
- [x] making simple GET/POST requests
- [x] learning how to call linux shell commands from golang
- [x] learning how to test golang in CI pipeline
- [x] learning how to build binaries/docker images for golang in CI pipeline