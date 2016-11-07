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
apt-get install -y python python-pip salt-master salt-minion salt-api python-cherrypy3
# pip install --upgrade pip cherrypy
# salt-call --local tls.create_self_signed_cert
mkdir -p /etc/pki/api/certs
openssl genrsa -out /etc/pki/api/certs/server.key 4096
openssl req -new -x509 -days 36500 -key /etc/pki/api/certs/server.key -out /etc/pki/api/certs/server.crt -subj "/CN=api.salt.com"

cat > /etc/salt/master.d/api.conf <<EOF
#rest_tornado:
rest_cherrypy:
  port: 8000
  ssl_crt: /etc/pki/api/certs/server.crt
  ssl_key: /etc/pki/api/certs/server.key
#  debug: True
#  disable_ssl: False
#  websockets: True
EOF
cat > /etc/salt/master.d/eauth.conf <<EOF
external_auth:
  pam:
    salt:
      - .*
      - '@wheel'
      - '@runner'
EOF

useradd -M -s /sbin/nologin salt 2> /dev/null
echo salt:salt | chpasswd
cat > /etc/salt/minion.d/minion.conf <<EOF
master: 192.168.88.101
EOF
echo master > /etc/salt/minion_id

#rsync
cat > /etc/rsyncd.conf <<EOF
uid=root
gid=root
max connection = 300
pid file = /var/run/rsyncd.pid
lock file = /var/run/rsyncd.lock
log file =/var/log/rsyncd.log

[ubuntu]
	comment = data ddb/stat
	path = /home/ubuntu/
	hosts allow = 100.0.0.0/8 10.0.0.0/8 127.0.0.1/16 192.168.0.0/16
	read only = yes
	timeout = 600
EOF

/usr/bin/rsync --daemon --config=/etc/rsyncd.conf
systemctl restart salt-minion
systemctl restart salt-master
systemctl restart salt-api
