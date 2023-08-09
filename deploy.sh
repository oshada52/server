#!/bin/bash
sshcmd="ssh -t oshada@ceyloninfo.site"
$sshcmd screen -S "deployment" /home/oshada/server/prod_deploy.sh