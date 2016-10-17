# usage

```
vagrant up
vagrant ssh master
```

## accept all keys

```
sudo salt-key -A -y
```

## test.ping

```
sudo salt '*' test.ping
```

## api

login

```
curl -k https://192.168.88.101:8000/login -H "Accept: application/x-yaml" -d username='salt' -d password='salt' -d eauth='pam'
```

test.ping

```
curl -k https://192.168.88.101:8000 -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab" -d client=local -d tgt='*' -d fun='test.ping'
```

request

```
curl -k https://192.168.88.101:8000/minions/minion -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab"
```

jobs list

```
curl -k https://192.168.88.101:8000/jobs -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab"
```


