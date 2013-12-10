# Beego Web

[![Requirement >= Go 1.2rc1](http://b.repl.ca/v1/Requirement-%3E%3D_Go%201.2rc1-blue.png)]() [![Requirement >= beego 0.9.9](http://b.repl.ca/v1/Requirement-%3E%3D_beego%200.9.9-blue.png)]()

An open source project for official documentation website of beego app framework.

## Install site locally

Beego Web is a `go get` able project:

	$ go get github.com/beego/beeweb

Switch to project root path:

	$ cd $GOPATH/src/github.com/beego/beeweb

Build and run with Go tools:

	$ go build
	$ ./beeweb

Or build with bee tool:

	$ bee run beeweb

Open your browser and visit [http://localhost:8090](http://localhost:8090).

## Build as your site

This project can be easily transferred as your own documentation site, there are some tips that you may want to know:

- In the file `conf/app.ini`:
	
	- `lang -> types`: languages that you want to support
	- `lang -> names`: user-friendly name of languages.
	- `app -> doc_names`: sections' name of additional documentation list.
	- It's **NOT** necessary but if you want to you can use GitHub app keys as following format:
		
			[github]
			client_id=1862bcb2******f36c
			client_secret=308d71ab53ccd858416cfceaed52******53c5f

- In the file `conf/navTree.json`:

	- This file is for design navigate bar and relation between sections and nodes.
	- After you modified this file, you need to change i18n setting in `conf/locale_<lang>.ini` files in order to show name correctly.
- In the file `conf/docTree.json`:

	- This file saves the file tree(with file name and commit) of your project that is hosted in GitHub. About how to use documentation project please see [beedoc](http://github.com/beego/beedoc). Note that if you added new section to documentation list and you do not want to wait auto-refresh, simple delete this file and restart.
	- To change the documentation project URL, you need to change it in function `checkDocUpdates` in file `models/models.go`.