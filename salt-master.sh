#! /usr/bin/env bash
cat > /etc/apt/sources.list <<EOF
deb http://cn.archive.ubuntu.com/ubuntu/ xenial main restricted universe multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-security main restricted universe multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-updates main restricted universe multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-proposed main restricted universe multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ xenial-backports main restricted universe multiverse

deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial main restricted universe multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-security main restricted universe multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-updates main restricted universe multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-proposed main restricted universe multiverse
deb-src http://cn.archive.ubuntu.com/ubuntu/ xenial-backports main restricted universe multiverse

deb http://repo.saltstack.com/apt/ubuntu/16.04/amd64/latest xenial main
EOF
# export http{,s}_proxy=http://10.94.97.161:8080/
wget -O - https://repo.saltstack.com/apt/ubuntu/16.04/amd64/latest/SALTSTACK-GPG-KEY.pub | sudo apt-key add -
apt-get update
apt-get install -y salt-master salt-minion salt-api salt-ssh
mkdir -p /etc/pki/tls/certs
openssl genrsa -out /etc/pki/tls/certs/localhost.key 4096
openssl req -new -x509 -days 36500 -key /etc/pki/tls/certs/localhost.key -out /etc/pki/tls/certs/localhost.crt \
        -subj "/C=CN/ST=beijing/L=shuzishangu/O=didichuxing/OU=Salt Api Root CA"
cat > /etc/salt/master.d/api.conf <<EOF
rest_cherrypy:
  port: 8000
  ssl_crt: /etc/pki/tls/certs/localhost.crt
  ssl_key: /etc/pki/tls/certs/localhost.key
  debug: True
EOF
cat > /etc/salt/master.d/eauth.conf <<EOF
external_auth:
  pam:
    salt:
      - .*
      - '@wheel'
      - '@runner'
EOF

useradd -M -s /sbin/nologin -p salt salt
cat > /etc/salt/minion.d/minion.conf <<EOF
master: 192.168.88.101
EOF
echo master > /etc/salt/minion_id

systemctl restart salt-minion
systemctl restart salt-master
systemctl restart salt-api
