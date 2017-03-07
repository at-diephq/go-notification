#!/usr/bin/env bash
if [ -n "$SSH_CLIENT" ];then 
    LOGIN_NOTIFY_HOST="localhost:8000"
    REMOTE_IP=`echo $SSH_CLIENT|awk '{print $1}'`
    LOGIN_NOTIFY_API="http://${LOGIN_NOTIFY_HOST}/notify/login?user=${USER}&remoteip=$REMOTE_IP"
    /usr/bin/curl "$LOGIN_NOTIFY_API" &> /dev/null
fi
