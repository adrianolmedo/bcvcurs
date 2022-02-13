# Venezolan currencies

RESTful API to obtain values of official currencies in Venezuela.

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

### **Get all currencies**

**GET:** `/`

Response (200 Ok):

```json

```

---

### **Get dollar**

**GET:** `/dollar`

Response (200 Ok):

```json

```
