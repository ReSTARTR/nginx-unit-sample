Nginx UNIT sample
====

NOTE:
This repository depends on the repository http://github.com/ReSTARTR/unit that is forked from https://github.com/nginx/unit

Build Nginx Unit from source
----

Install dependencies(ref: https://github.com/nginx/unit#precompiled-packages).

Build unitd control server

```sh
$ go get github.com/ReSTARTR/unit # or git-clone
$ cd github.com/ReSTARTR/unit
$ git checkout go-include-path
$ ./configure && make
```

Run unitd

```sh
$ ./build/unitd --path=/tmp/control.unit.sock
```

Prepare for each app server

```sh
for lang in php python go; do
  ./configure $lang && make $lang
done
```


Build Go Server
----

```sh
$ cd ./go_server
$ go build -o go_server
```

Notify to unit control server.

```sh
$ cd github.com/ReSTARTR/unit
$ curl --unix /tmp/control.unit.sock -XPUT -d @./listeners.json http://localhost/
```

and request to the application server

```sh
$ for port in 8200 8300 8400; do curl localhost:$port; echo; done
Hello, PHP
Hello, Python
Hello, Go
```
