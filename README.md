Description
===================
yaml2json


Build
===================
Build & install into `$GOPATH/bin`:
```
go get -u github.com/dolmen/yaml2json
go install github.com/dolmen/yaml2json
```

Usage
====================
### shell
* run `echo "a: 1" | ./yaml2json` to see result

### read from file, save to file
```
yaml2json < in.json > out.json
```

References
====================
* https://github.com/bronze1man/yaml2json
* https://github.com/peter-edge/go-yaml2json
