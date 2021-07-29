# Learn API security (in Golang)

## Prequisite

- GO 1.16 or later
- Mysql
- curl

### Mysql on Docker

```
docker run --name go-apisecurity -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql
```

## What we learn

### Injection Attack

- SQL Injection
    - [x] Vuln
    - [x] Fix
- XSS Attack
    - [ ] Vuln
    - [ ] Fix

### Securing API

- [x] Rate Limiting
- Authentication
- Encryption
- Audit Logging
- Access Control