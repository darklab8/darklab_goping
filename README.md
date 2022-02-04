golang program to query same url N times at the same time

example:

$: go run . 125 https://my-secret-url.com/ping
```
{"message":"pong!"}%!(EXTRA time.Duration=585.9592ms)
{"message":"pong!"}%!(EXTRA time.Duration=673.824854ms)
{"message":"pong!"}%!(EXTRA time.Duration=877.372085ms)
...
{"message":"pong!"}%!(EXTRA time.Duration=865.47524ms)
{"message":"pong!"}%!(EXTRA time.Duration=813.026781ms)
total time 1.055035429s
```