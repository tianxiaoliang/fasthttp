
First build the fuzz binary:
```bash
$ go-fuzz-build -o fuzz.zip github.com/erikdubbelboer/fasthttp/fuzz
```

Then run the fuzzer:
```bash
$ go-fuzz -bin=fuzz.zip -workdir=. -procs=6 -dup
```

