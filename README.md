# supreme-dollop

* Tasks: https://trello.com/b/qFKy7gUY/dogr
* Design Docs: https://drive.google.com/drive/folders/1IHg6UDWjfwaa4AXUuFvRc06TREEDsVu-

## Installation instructions
TBD
If using vscode, just use the devcontainer

### GO SERVER
To build and run go server
```
cd go-api
go build main.go
./main
```

Control + C to kill server. 

If you accidentally used Control + Z, use
```
jobs
kill %1
```

### TYPESCRIPT FRONT END
Built using Node v16.9.0

* To install dependencies: npm install 
* To build and run development: npm run dev
* To build for production: npm run build
* To preview: npm run preview

== VERSIONS ==
* Go 1.19
* Typescript 4.9.5
* Node 16.9.0