# go-rest

modified version of [go-rest](https://github.com/r3labs/go-salt)

detailed api reference see [rest_cherrypy](https://docs.saltstack.com/en/latest/ref/netapi/all/salt.netapi.rest_cherrypy.html)

## feature

- retry authenticate after receive 401
- easy access to urls /jobs, /minions and /run

## usage

- run test.ping

```
client := NewClient(addr, username, password, true)
ping := map[string]bool{}
err := client.RunCmd("*", "test.ping", nil, &ping)
fmt.Print(ping)
```

# vagrant usage

```
vagrant up
vagrant ssh master
```

## command via shell


- accept all keys

```
sudo salt-key -A -y
```
- ping

```
sudo salt '*' test.ping
```

## api via curl

- login

```
curl -k https://192.168.88.101:8000/login -H "Accept: application/x-yaml" -d username='salt' -d password='salt' -d eauth='pam'
```

- test.ping

```
curl -k https://192.168.88.101:8000 -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab" -d client=local -d tgt='*' -d fun='test.ping'
```

- request

```
curl -k https://192.168.88.101:8000/minions/minion -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab"
```

- jobs list

```
curl -k https://192.168.88.101:8000/jobs -H "Accept: application/x-yaml" -H "X-Auth-Token: b04cb118d4a0db2213b09db2f42c42ab"
```


