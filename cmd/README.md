# salt cmd package

## feature

- fullly realize of salt cmd model
- interface see [cmd.go](cmd.go)

## usage

see document on salt
```
sudo salt '*' sys.doc cmd.run
```

cmd.run
```
addr     = "192.168.88.101:8000"
username = "salt"
password = "salt"
target = "minion"
client = salt.New(target, salt.NewClient(addr, username, password, true))
r, err := client.Run("whoami", &Param{Runas: "ubuntu"})
```

