#!/bin/sh

echo $PWD
mkdir -p $PWD/data

./menu create --artifact-id=menu-webapp --artifact-version=1.0.0 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip --cookbook=https://github.com/ngerakines/menu-webapp-cookbook file://$PWD/data/`date +%s`1.menu
./menu create --artifact-id=menu-webapp --artifact-version=1.0.1 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip --cookbook=https://github.com/ngerakines/menu-webapp-cookbook file://$PWD/data/`date +%s`2.menu
./menu create --artifact-id=menu-webapp --artifact-version=1.2.0 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip --cookbook=https://github.com/ngerakines/menu-webapp-cookbook file://$PWD/data/`date +%s`3.menu
./menu create --artifact-id=menu-webapp --artifact-version=1.2.1 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip --cookbook=https://github.com/ngerakines/menu-webapp-cookbook file://$PWD/data/`date +%s`4.menu
