## Contributing

# CombatIntel

CombatIntel is a backend system for tactical analysis and prediction of combat missions.  
Built in **Go** with clean architecture, **PostgreSQL**, and **sqlc**.  
Supports JWT authentication using **RSA key pairs**, and machine learning predictions.

---

## Features

- Role-based access: `admin`, `officer`
- JWT auth with RSA (`private_key.pem`, `public_key.pem`)
- CRUD for:
    - Units
    - Users
    - Missions (includes outcome, losses, notes, CSV upload)
- Prediction model training (WIP)

---

## .env Example

```env
DATABASE_URL=
TEST_DATABASE_URL=
PORT="8080"
```

---

JWT RSA Setup
Generate RSA key pair:

These are used for signing (private) and verifying (public) JWT tokens with RS256 algorithm.
```bash
openssl genrsa -out keys/private_key.pem 2048
openssl rsa -in keys/private_pem.key -pubout -out keys/public_key.pem
```

## Requirements
- Go 1.22+
- PostgreSQL 14+
- sqlc
- RSA keys (keys/private_pem.key, keys/public_pem.key)

