
## platform service (from dev)

### 发email给某些用户


```
POST /push?mode=<mode>&id[]=<userId>&id[]=<userId2>...
Authorization: appName

<messageBody>
```

* `<mode>`为0时表示单独给每个收件人发邮件。此为默认值
* `<mode>`为1时表示所有人都在收件列表内

### 发email给所有这个app订阅用户


```
POST /pushall?mode=<mode>
Authorization: appName

<messageBody>
```

* `<mode>`为0时表示单独给每个收件人发邮件。此为默认值
* `<mode>`为1时表示所有人都在收件列表内

## hook (to dev)

```
POST <appHost>/push?id=<userId>

<messageBody>
```

```
POST <appHost>/nofication

id=<userId>&action=subscribe
```


## 内部(from webservice)

```
POST /user/subscribe

id=<userId>&appName=<appName>
```
