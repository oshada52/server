#!/bin/bash
SECONDS=0

cd $HOME/server

msg() {
  echo -e "$1\n------------------\n"
}


msg "Pulling from Github"
git pull

msg "Building Docker image"
sudo docker build --tag server .

msg "Stopping Docker container"
sudo docker stop server_c
sudo docker rm server_c

msg "Starting Docker container"
sudo docker run -d --name server_c --expose 443 -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt -e SERVER_ENV=PROD server

duration=$SECONDS

echo
msg "Deploy finished in $(($duration % 60)) seconds."
msg "Press Enter to exit"
read