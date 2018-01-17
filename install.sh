#!/bin/bash

scp -i ~/.ssh/CryptoBookApp.pem bin/start.sh ec2-user@ec2-18-218-97-175.us-east-2.compute.amazonaws.com:~/saithal/sandbox/cryptoapp-data-collection/bin/
scp -i ~/.ssh/CryptoBookApp.pem bin/stop.sh ec2-user@ec2-18-218-97-175.us-east-2.compute.amazonaws.com:~/saithal/sandbox/cryptoapp-data-collection/bin/
scp -i ~/.ssh/CryptoBookApp.pem dev.config.yaml ec2-user@ec2-18-218-97-175.us-east-2.compute.amazonaws.com:~/saithal/sandbox/cryptoapp-data-collection/
scp -i ~/.ssh/CryptoBookApp.pem crypto-data-collection ec2-user@ec2-18-218-97-175.us-east-2.compute.amazonaws.com:~/saithal/sandbox/cryptoapp-data-collection/bin/