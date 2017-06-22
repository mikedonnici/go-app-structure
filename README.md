## Sample structures for a Go app

**1. Top-level main**

```
project/
├── main.go # primary executable
├── cmd/ # contains code used for cli commands
│   ├── executable1/ # executables require package main so have 'main.go'
│   │	└── main.go
│   └── executable2/
│       └── main.go
├── internal/ # contains internal packages not useful anywhere else
│   ├── pkg1/ # primary package file named after package (dir)
│   │   └── pkg1.go
│   └── pkg2/
│       └── pkg2.go
└── vendor/ # third party packages managed with govendor - see below
    └── vendor.json
```

**2. Side-by-side version**

```
project/
├── app/ # contains main app code - could be 'server'?
│   └── main.go
├── cmd/ # contains code used for cli commands
│   ├── executable1/ # executables require package main so have 'main.go'
│   │	└── main.go
│   └── executable2/
│       └── main.go
├── internal/ # contains internal packages not useful anywhere else
│   ├── pkg1/ # primary package file named after package (dir)
│   │   └── pkg1.go
│   └── pkg2/
│       └── pkg2.go
└── vendor/ # third party packages managed with govendor - see below
    └── vendor.json
``` 

###Heroku Deployment 

To deploy the examples above requires a dependency manager such as [govendor](https://github.com/kardianos/govendor). The `"heroku` field in `vendor.json` specifies which executables should be built. The Heroku `Procfile` is used to specify the executables.
  
**Example 1**
 
`vendor.json` 

```json
{
  "rootPath": "github.com/mikedonnici/go-app-structure",
  "heroku": {
		"goVersion": "go1.8",
		"install": [
            ".",
			"./cmd/..."
		]
	}
}
```

`Procfile`

```text
app: go-app-structure
hi: hello
bye: goodbye
```

**Example 2**

`vendor.json`

```json
{
  "rootPath": "github.com/mikedonnici/go-app-structure",
  "heroku": {
		"goVersion": "go1.8",
		"install": [
            "./app/...",
			"./cmd/..."
		]
	}
}
```

`Procfile`

```text
app: app
hi: hello
bye: goodbye
```
 
The process type names (on the left) are arbitrary, except for `web`, which is reserved for the only process that can receive http requests.

The executables can be called from the Heroku cli thus:

```bash
$ heroku run app
$ heroku run hi
$ heroku run bye

# Can also run them by the name of the executable
$ heroku run go-app-structure
$ heroku run hello
$ heroku run goodbye
```


Example for MappCPD adapted from [here](https://www.goinggo.net/2017/02/package-oriented-design.html)

```
github.com/mappcpd/api
├── cmd/
│   ├── servi/
│   │   ├── cmdupdate/
│   │   ├── cmdquery/
│   │   └── servi.go
│   └── servid/
│       ├── routes/
│       │   └── handlers/
│       ├── tests/
│       └── servid.go
├── internal/
│   ├── attachments/
│   ├── locations/
│   ├── orders/
│   │   ├── customers/
│   │   ├── items/
│   │   ├── tags/
│   │   └── orders.go
│   ├── registrations/
│   └── platform/
│       ├── crypto/
│       ├── mongo/
│       └── json/
└── vendor/
    ├── github.com/
    │   ├── ardanlabs/
    │   ├── golang/
    │   ├── prometheus/
    └── golang.org/
```




### References

https://golang.org/doc/code.html

https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/packaging/README.md

https://www.goinggo.net/2017/02/package-oriented-design.html
 
 
 


 
 
 
 
