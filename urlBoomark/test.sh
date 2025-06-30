#!/bin/sh

curl -X POST -H "application/json" -d '{"Name":"Facebook","Url":"https://facebook.com"}' http://localhost:8080/bookmarks
