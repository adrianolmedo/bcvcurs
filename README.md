# Venezolan currencies

RESTful API to obtain values of official currencies in Venezuela directly from BCV (http://www.bcv.org.ve/).

## Run local

```bash
$ git clone https://github.com/adrianolmedo/vecurs.git
$ go install .
$ vecurs -addr localhost -port 8080
```

## Run with Docker

```bash
$ git clone https://github.com/adrianolmedo/vecurs.git
$ docker build --tag vecurs:0.1 .
$ docker run -d -p 8080:80 --name vecurs vecurs:0.1
```

## Endpoints

### **All currencies**

**GET:** `/v1`

Response (200 Ok):

```json
{
    "data": [
        {
            "iso": "EUR",
            "symbol": "€",
            "value": 5.04244811
        },
        {
            "iso": "CNY",
            "symbol": "¥",
            "value": 0.7005779
        },
        {
            "iso": "TRY",
            "symbol": "₺",
            "value": 0.3255186
        },
        {
            "iso": "RUB",
            "symbol": "₽",
            "value": 0.05899348
        },
        {
            "iso": "USD",
            "symbol": "$",
            "value": 4.4369
        }
    ],
    "message_ok": {
        "content": ""
    }
}

```

---

### **Euro**

**GET:** `/euro`

Response (200 Ok):

```json
{
    "data": {
        "iso": "EUR",
        "symbol": "€",
        "value": 5.04244811
    },
    "message_ok": {
        "content": ""
    }
}
```

---

### **Yuan**

**GET:** `/yuan`

Response (200 Ok):

```json
{
    "data": {
        "iso": "CNY",
        "symbol": "¥",
        "value": 0.7005779
    },
    "message_ok": {
        "content": ""
    }
}
```

---

### **Lira**

**GET:** `/lira`

Response (200 Ok):

```json
{
    "data": {
        "iso": "TRY",
        "symbol": "₺",
        "value": 0.3255186
    },
    "message_ok": {
        "content": ""
    }
}
```

---

### **Ruble**

**GET:** `/ruble`

Response (200 Ok):

```json
{
    "data": {
        "iso": "RUB",
        "symbol": "₽",
        "value": 0.05899348
    },
    "message_ok": {
        "content": ""
    }
}
```

---

### **Dollar**

**GET:** `/dollar`

Response (200 Ok):

```json
{
    "data": {
        "iso": "USD",
        "symbol": "$",
        "value": 4.4369
    },
    "message_ok": {
        "content": ""
    }
}
```
