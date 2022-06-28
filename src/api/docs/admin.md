### 用户登陆

目前是用账户密码登陆，登陆账号密码为ppai，ppapi

#### Request

- Method POST
- URI /admin/v1/login
- Body

```json5
{
  "username": "ppapi",
  "password": "ppapi"
}
```

#### Response

- Body

```json
{
  "succeeded": true,
  "resp_data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY0MzMwMjAsIlVpZCI6MSwiUm9sZSI6Mn0.RzQE4mh8tqJE1eBRR1bpiwB6cpWbjPyVE5gTRutRPSY",
    "admin_info": {
      "id": 1,
      "username": "ppapi",
      "password": "836ed0d4c9fc8a9f9484f8b004c6f720",
      "nickname": "ppapi",
      "created_at": 1645033731,
      "updated_at": 1645033888
    }
  }
}
```