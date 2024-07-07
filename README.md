GPT API
=======

Basic golang experiments with GPT models.

Purpose
-------

Repo was created to learn some basics of Go language.

Setup
-----

```shell
cat .env.example > .env
docker compose build
docker compose up
```

Edit `.env` to set proper values for your `env` variables.

Docs
----

Use `swag` package to generate API docs:

```shell
cd backend
swag init
```