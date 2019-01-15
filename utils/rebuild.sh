#!/bin/sh

echo Building new version...

cd ../src
go build
mv src ../bin/nsd

echo All done!