#! /bin/sh

kill -9 $(pgrep filestore-server)
cd /opt/golang/src/filestore-server
git pull https://github.com/laihezhao/filestore-server.git
chmod +x filestore-server
./filestore-server