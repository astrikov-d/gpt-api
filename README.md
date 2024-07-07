GPT API
=======

Basic golang experiments with GPT models.

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