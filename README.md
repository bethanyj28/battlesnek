# battlesnek
My API for [Battlesnake](https://play.battlesnake.com/).

## Getting Started
To run this api locally, simply run `make build && make run`, and it will launch the server on `localhost:8080`. It follows the API defined in the [Battlesnake docs](https://docs.battlesnake.com/references/api#the-battlesnake-api). There is an additional `/health` endpoint.

## Structure

### cmd
`cmd/server` defines the server functionality for the server. This is where the handler and routes are defined.

### internal
`internal` contains general functionality that can be abstracted for the program. This includes structs, errors, and packages.

`battle` contains different implementations of the `snake` interface.

## TODO
- [x] Deployments with Github Actions and Heroku
- [x] Set up base snake API
- [x] Create a more intelligent snake
- [ ] Dynamically switch snakes with authenticated endpoint
- [x] Add testing suite
- [x] Prioritize directions
