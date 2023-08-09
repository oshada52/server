#!/bin/bash
SECONDS=0

cd $HOME/server

msg() {
  echo -e "$1\n------------------\n"
}

msg "Stopping app"
sudo pkill server

msg "Pulling from Github"
git pull


msg "Building Go binary"
go build

msg "Starting server"
nohup sudo ./server &/dev/null &

duration=$SECONDS

echo
msg "Deploy finished in $(($duration % 60)) seconds."
msg "Press Enter to exit"
read