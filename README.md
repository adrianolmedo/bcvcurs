# BCV currencies

RESTful API to obtain values of official currencies in Venezuela directly from BCV (http://www.bcv.org.ve/).

## Run local

```bash
$ git clone https://github.com/adrianolmedo/bcvcurs.git
$ go install .
$ bcvcurs -addr localhost -port 8080
```

## Run with Docker

```bash
$ git clone https://github.com/adrianolmedo/bcvcurs.git
$ make
```

Note: `make` by default runs the target defined in the `.DEFAULT_GOAL` variable of the `Makefile`.

## Content

* [Endpoints](#endpoints)
    * [All currencies](#all-currencies)
    * [Euro](#euro)
    * [Yuan](#yuan)
    * [Lira](#lira)
    * [Ruble](#ruble)
    * [Dollar](#dollar)
* [Errors](#errors)
  * [Error path](#error-path)
  * [Getting data](#getting-data)

## Endpoints

### **All currencies**

**GET:** `/v1`

Response (200 OK):

```json
{
    "data": {
        "dollar": {
            "iso": "USD",
            "symbol": "$",
            "value": 4.4032
        },
        "euro": {
            "iso": "EUR",
            "symbol": "€",
            "value": 5.00608614
        },
        "lira": {
            "iso": "TRY",
            "symbol": "₺",
            "value": 0.3234693
        },
        "ruble": {
            "iso": "RUB",
            "symbol": "₽",
            "value": 0.05774612
        },
        "yuan": {
            "iso": "CNY",
            "symbol": "¥",
            "value": 0.694785
        }
    },
    "message_ok": {
        "content": ""
    }
}
```

---

### **Euro**

**GET:** `/euro`

Response (200 OK):

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

Response (200 OK):

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

Response (200 OK):

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

Response (200 OK):

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

Response (200 OK):

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

---

## **Errors**

### **Error path**

Response (404 Not Found):

```json
{
    "message_error": {
        "content": "path error"
    }
}
```

### **Getting data**

Response (503 Service Unavailable):

```json
{
    "message_error": {
        "content": "error getting data"
    }
}
```