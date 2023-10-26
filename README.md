# Go Bank App With Go

This app is used to transfer amount of money from one account to another account.

## How to Run

Type on your VSCode:
```bash
go run main.go
```

## How to Test

POST /register for register a user
```bash
{
    "username": "Sanji",
    "email": "onepiece@mail.com",
    "password": "Sanji123"
}
```
```bash
{
    "data": {
        "ID":3,
        "Username":"Sanji",
        "Email":"onepiece@mail.com",
        "Accounts": [
            {
                "ID":3,
                "Name":"Sanji's account",
                "Balance":0
            }
        ]
    },
    "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzA2MTUsInVzZXJfaWQiOjN9.t5OAc_tgD1NpoYf0a13Llsoq3sRV39u4k-b1rejPZHg",
    "message":"all is fine"
}
```


POST /login for login a user
```bash
{
    "username": "Sanji",
    "password": "Sanji123"
}
```
```bash
{
    "data": {
        "ID":3,
        "Username":"Sanji",
        "Email":"onepiece@mail.com",
        "Accounts": [
              {
                "ID":3,
                "Name":"Sanji's account",
                "Balance":0
              }
        ]
    },
    "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzg2MTksInVzZXJfaWQiOjN9.9frlJM6Bcnu0mktyWlIlKibNLd1eGu7_ni71JbGdmh4",
    "message":"all is fine"
}
```


GET /user/{id} for get detail of a user by Id

To get user by Id enter the bearer token for authorization
```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzA2MTUsInVzZXJfaWQiOjN9.t5OAc_tgD1NpoYf0a13Llsoq3sRV39u4k-b1rejPZHg

{
    "data": {
        "ID":3,
        "Username":"Sanji",
        "Email":"onepiece@mail.com",
        "Accounts": [
              {
                  "ID":3,
                  "Name":"Sanji's account",
                  "Balance":0
              }
        ]
    },
    "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzg2MTksInVzZXJfaWQiOjN9.9frlJM6Bcnu0mktyWlIlKibNLd1eGu7_ni71JbGdmh4",
    "message":"all is fine"
}
```


POST /transaction for create transaction between user

To do transaction, user need to login first then enter the bearer token for authorization
```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzY0ODAsInVzZXJfaWQiOjF9.oTYit1OSQKlL2fv84C0aCIPLl_Tp7qAOajMUAZNDKWE

{
    "UserId": 1,
    "From": 1,
    "To": 2, 
    "Amount": 200000
}
```
```bash
{
    "data":
        {"ID":1,
        "Name":"saiful's account",
        "Balance":4800000},
        "message":"all is fine"
}
```


GET /transaction/{id} for get history of transaction

To get history of transaction, user need to login first then enter the bearer token for authorization
```bash
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2OTgzMzY0ODAsInVzZXJfaWQiOjF9.oTYit1OSQKlL2fv84C0aCIPLl_Tp7qAOajMUAZNDKWE

{
    "data": [
        {
            "ID":1,
            "From":1,
            "To":2,
            "Amount":200000
        },
        {
            "ID":2,
            "From":1,
            "To":2,
            "Amount":500000
        }
    ],
    "message":"all is fine"
}
```
