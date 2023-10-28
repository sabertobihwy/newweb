#! /bin/sh
kill -9 $(pgrep webserver)
cd ~/newweb/
git pull git@github.com:sabertobihwy/newweb.git
cd webserver/
./webserver &
