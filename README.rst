Go File Server
==============

This program serves files in a directory.

Usage
-----

::

  gofileserver -port [port] -dir [directory]

If directory is not specified, the current working
directory will be served.

If port number is not specified, 9999 is used
as the default port number.

The file server can be accessed using all available
IPs or hostnames in the system.

Building from Source
--------------------

Setup the Go development environment as documented here.
http://baijum.blogspot.in/2013/09/go-programming-installation.html

After setting up go environment clone this repository::

  git clone git@github.com:baijum/gofileserver.git

Run `go build` inside gofileserver::

  cd gofileserver
  go build gofileserver.go

Copy the `gofileserver` to /usr/local/bin::

  sudo cp gofileserver /usr/local/bin

Run the application as given the Usage section.
