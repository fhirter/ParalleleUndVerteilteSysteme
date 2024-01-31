#!/usr/bin/env bash

# to run supply ip: ./setup.sh 192.12.11.4

IP=$1
scp backup.sh root@$IP:/root

ssh root@$IP << "EOF"

apt update

apt install -y postgresql postgresql-contrib apache2 ufw restic
systemctl start postgresql.service
systemctl enable apache2
systemctl start apache2

ufw allow OpenSSH
ufw allow 'Apache Full'
yes | ufw enable

mkdir /root/data
for i in {1..1000}; do base64 /dev/urandom | head -c 1000000 > "/root/data/dummyfile$i.txt"; done

(crontab -l 2>/dev/null; echo "* * * * * /root/backup.sh") | crontab -

EOF