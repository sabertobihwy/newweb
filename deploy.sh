#! /bin/sh
kill -9 $(pgrep webserver)
cd $(pwd)/newweb/webserver/
git pull git@github.com:sabertobihwy/newweb.git
cd webserver/
./webserver &
