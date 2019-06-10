# Planet Users
Implements a simple user/group management rest API

## Installation
- Install Docker
- Install docker-compose
- Clone repository locally
- you do not have to have any go tools installed, but it will make your life easier to inspect the code

## Execution
- Run `docker-compose up`

## Architecture
### Database
The docker compose service `db` will initialize a mysql database 
and create a tiny seed database as described in `sql/bootstrap.sql`
### Web
The docker compose service `web` will run a web server on port `8080` and expose
it on your local machine.  if this port is in use, docker compose will give
you and error and you can change the local port binding to a different port.


## Notes

- I used a familiar package `https://github.com/volatiletech/sqlboiler` to handle persistance against the mysql database
- I used chi for http routing and logrus, which are very common packages
- The code in `dbmodels` is generated
 


