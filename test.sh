#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{"url":"https://myspace.com"}' http://localhost:8080/short
