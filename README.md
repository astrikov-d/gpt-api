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

Open http://localhost/ping to check that you'll get proper response like

```json
{
  "message": "pong"
}
```