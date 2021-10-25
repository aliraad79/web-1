## Contribution
- Jafar Sadeghi 97106079
- Ali Ahmad 97105703
- Hamed Abdi 96109782

## Installation
All files and requirements can be installed via docker. so just install [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/) and then run

```sh
    docker-compose up -d
```
in terminal.


### Test
run 
```sh
    docker-compose up -d --scale worker=4
```
open `http://localhost:8089` and start test with host `http://ngingx`
then run 
```sh
    docker-compose up --scale worker=4 --scale go=3 --scale node=3
```
and try it again
