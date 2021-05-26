# xtorrent backend API

## User - Related

### user register

`POST /register`

body:

```json
{
  "email_address": ""
}
```
return:

- if error occur:
```json
{
  "code": 400,
  "message": ""
}
```

- success:
```json
{
  "code": 200,
  "message": "ok",
  "token": ""
}
```

### User login


`POST /login`

body:

```json
{
  "email_address": "",
  "password": ""
}
```
return:

- if error occur:
```json
{
  "code": 400,
  "message": ""
}
```

- success:
```json
{
  "code": 200,
  "message": "ok",
  "token": ""
}
```

### Refresh JWT token


`POST /token`

header:
```
Authorization: Berear ${JWT TOKEN}
```

return: 

- if error occur:
```json
{
  "code": 400,
  "message": ""
}
```

- success:
```json
{
  "code": 200,
  "message": "ok",
  "token": ""
}
```

### Torrent - Related

### Add torrent

`POST /register`

header:

```
Authorization: Berear ${JWT token}
```

body:

```json
{
  "magnet": "",
  "picture": file,
  "name": "",
  "description": ""
}
```

return:

```json
{
  "code": 400,
  "message": ""
}
```

### Get torrent

```
GET /torrent?insertStartTime=${int64}
             number=${int}
             name=""
```

return:

- if error occur

```json
{
  "code": 400,
  "message": "${error message}"
}
```

- success

```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "number": 10,
    "torrent_infos": [
      {
        "magnet": "",
        "picture_path": ["", ""],
        "name": "",
        "description": "",
        "insert_time" : `${unixTimeStamp}`
      }
    ]
  }
}
```

## Control - Related

### Refresh config

`PUT /ctrl/config`

> re-read config from config file


### Tool chain

- go echo web framework  -> v4
- jWT middleware -> for authentication