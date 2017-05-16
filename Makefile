#! /usr/bin/make
#
# Makefile for goa cellar example
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
# - ae-build  build appengine
# - ae-dev    deploy to local (dev) appengine
# - ae-deploy deploy to appengine
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#

all: depend clean generate build

depend:
	@glide install

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf public/swagger
	@rm -rf public/schema
	@rm -rf public/js
	@rm -f todo

bootstrap:
	@goagen main    -d github.com/m0a-mystudy/goa-todo/design -o controllers

generate:
	@goagen app     -d github.com/m0a-mystudy/goa-todo/design
	@goagen swagger -d github.com/m0a-mystudy/goa-todo/design -o public
	@goagen schema  -d github.com/m0a-mystudy/goa-todo/design -o public
	@goagen client  -d github.com/m0a-mystudy/goa-todo/design
	@goagen js      -d github.com/m0a-mystudy/goa-todo/design -o public

build:
	@go build -o todo

# ae-build:
# 	@if [ ! -d $(HOME)/cellar ]; then \
# 		mkdir $(HOME)/cellar; \
# 		ln -s $(CURRENT_DIR)/appengine.go $(HOME)/cellar/appengine.go; \
# 		ln -s $(CURRENT_DIR)/app.yaml     $(HOME)/cellar/app.yaml; \
# 	fi

# ae-deploy: ae-build
# 	cd $(HOME)/cellar
# 	python2 appcfg.py update .
