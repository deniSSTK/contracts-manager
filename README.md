# Project Setup

## Run

To start the project in development mode, follow these steps:

1. Copy the environment example file:

```bash
cp .env.example .env
```

2. (Optional) Update admin credentials in the `.env` file.

3. Start the project using Docker:

```bash
docker compose up -d
```

> This will start all services in the background and make them ready to use.

---

## Admin Credentials

The `.env` file contains default admin credentials that are used to access the system.
You can change them to whatever you want:

```env
ADMIN_USERNAME=admin_user
ADMIN_EMAIL=admin@email.com
ADMIN_PASSWORD=admin_pass
```

---

## Examples

The `./example` folder contains sample files for import:

* `.json` — for Contracts and Persons
* `.csv` — for Contracts and Persons

> Use them to quickly test the import functionality.
