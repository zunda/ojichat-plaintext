# ojichat-plaintext
Your friendly ojisan in the cloud inspired by [ojichat-web](https://reverent-shirley-368990.netlify.com/).

## Running and accessing the app
### Local
```
$ govendor sync
$ go build -o ojichat main.go
$ ./ojichat
```

Navigate to http://localhost:3000/

## As a Heroku app
```
$ heroku apps:create
$ git push heroku master
$ heroku apps:open
```

## In browser
```
$ GOOS=js GOARCH=wasm go build -o js/ojichat.wasm js.go
```

Publish the `js/` directory.

## License
- js/index.html: Copyright 2018 The Go Authors under [BSD-style](LICENSE-go)
- js/wasm_exec.js: Copyright 2018 The Go Authors under [BSD-style](LICENSE-go)
- Otherwise: Copyright 2019 zunda under [MIT](LICENSE)
