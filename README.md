## Sample structures for a Go app

**1. Top-level main**

```text
/project
 |
 |__ main.go # primary executable
 |
 |__/cmd # contains code used for cli commands
 |   |
 |   |__/executable1 # executables require package main so have 'main.go'
 |   |   |__ main.go
 |   |   |__/internal # packages only needed to executable1 
 |   |   |..
 |   |__/executable2
 |   |   |__ main.go
 |   |   |...
 |   |...
 |
 |__/internal # contains internal packages not useful anywhere else
 |   |
 |   |__/pkg1 # primary package file named after package (dir)
 |   |   |__ pk1.go
 |   |   |__/internal # packages only needed for pkg1
 |   |
 |   |__/pkg2
 |       |__ pkg2.go
 |
 |__/vendor # third party packages managed with govendor - see below
 |__ vendor.json
```

**2. Side-by-side version**

```text
/project
 |
 |__/app # contains main app code - could be 'server'?
 |   |
 |   |__/internal # contains internal packages not useful anywhere else
 |   |   |
 |   |   |__/pkg1 # primary package file named after package (dir)
 |   |   |   |__pk1.go
 |   |   |   |__/internal # packages only needed for pkg1
 |   |   |
 |   |   |__/pkg2
 |   |       |__pkg2.go
 |   |...
 |
 |__/cmd # contains code used for cli commands
 |   |
 |   |__/executable1 # executables require package main so have 'main.go'
 |   |   |__main.go
 |   |   |__/internal # packages only needed to executable1 
 |   |   |..
 |   |__/executable2
 |   |   |__main.go
 |   |   |...
 |   |...
 |
 |__/vendor # third party packages managed with govendor - see below
 |__ vendor.json
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

 
 
 


 
 
 
 