#!/bin/bash

APP_USER=ec2-user
APP=/home/ec2-user/saithal/sandbox/cryptoapp-data-collection/bin/crypto-data-collection
CONFIG=/home/ec2-user/saithal/sandbox/cryptoapp-data-collection/dev.config.yaml
LOGFILE=/home/ec2-user/saithal/sandbox/cryptoapp-data-collection/log/dev.log
PIDFILE=/tmp/cryptoapp-data-collection.pid

nohup $APP --config $CONFIG < /dev/null > $LOGFILE 2>&1 &
echo $! > $PIDFILE
