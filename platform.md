
## platform service (from dev)

### 发email给某些用户


```
POST /push?id[]=<userId>&id[]=<userId2>...
Authorization: appName

<messageBody>
```

### 发email给所有这个app订阅用户


```
POST /pushall
Authorization: appName

<messageBody>
```


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
