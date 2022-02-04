golang program to query same url N times at the same time.

example:

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