for running node individually install [node](https://nodejs.org/en/download/) then comment line 15 of index.js and uncomment line 16 after that

```sh
    docker run -p 6379:6379 -d redis:6-alpine
    yarn run run
```