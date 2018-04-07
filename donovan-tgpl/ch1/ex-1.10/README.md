# Http request cache #

## Steps to check ##

```bash
if [[ -f ".cache" ]]; then; rm .cache; fi;
go run main.go https://example.com # notice elapsed time
go run main.go https://example.com # compare elapsed time
```

## Expected output ##

```bash
go run main.go https://example.com
0.53s     1270  https://example.com
0.53s elapsed

go run main.go https://example.com
0.00s     1270  https://example.com
0.00s elapsed
```

## Known issues ##

This was a horrible experience.
