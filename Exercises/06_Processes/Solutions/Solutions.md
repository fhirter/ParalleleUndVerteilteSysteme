# Lösungen Processes

## Befehle
```shell
`Ps -el`
`Htop`
`top`
```

## Installation
```shell
apt update

apt install postgresql postgresql-contrib apache2 ufw restic
systemctl start postgresql.service
systemctl enable apache2
systemctl start apache2

ufw allow OpenSSH
ufw allow 'Apache Full'
ufw enable

ufw status
```