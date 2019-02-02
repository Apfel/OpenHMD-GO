#		Copyright (c) 2019 Marko (Apfel)
#
#		Permission is hereby granted, free of charge, to any person obtaining a copy
#		of this software and associated documentation files (the "Software"), to deal
#		in the Software without restriction, including without limitation the rights
#		to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
#		copies of the Software, and to permit persons to whom the Software is
#		furnished to do so, subject to the following conditions:
#
#		The above copyright notice and this permission notice shall be included in all
#		copies or substantial portions of the Software.
#
#		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#		IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#		FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
#		AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
#		LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
#		OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
#		SOFTWARE.
#
#		Fun fact: I copied this Makefile from my Discord-Bot.

#----------------------------------------------------------------------
# Go settings.
#----------------------------------------------------------------------
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get
GOTEST=$(GO) test

#----------------------------------------------------------------------
# Standard option to run.
#----------------------------------------------------------------------
all: build

#----------------------------------------------------------------------
# This should help some users.
#----------------------------------------------------------------------
help: 
	@echo "OpenHMD-GO - Makefile - Version 1.0-rc1"
	@echo " "
	@echo "make build 			- Builds OpenHMD. This is the default option."
	@echo "make help 			- Displays this message."
	@echo "make test 			- Tests the compiled OpenHMD module after building."
	@echo "make clean 			- Deletes all files created by building."
	@echo "make deps 			- Installs all dependencies"
	@echo "make depsremove 		- Removes all dependencies."

#----------------------------------------------------------------------
# General building procedure.
#----------------------------------------------------------------------
build:
	@echo "Building OpenHMD-GO..."
	@$(GOBUILD) -v

#----------------------------------------------------------------------
# Testing of the module.
#----------------------------------------------------------------------
test: build
	@echo "Initiating test of OpenHMD-GO..."
	@$(GOTEST)

#----------------------------------------------------------------------
# Cleanup.
#----------------------------------------------------------------------
clean:
	@echo "Cleaning directory..."
	@$(GOCLEAN)