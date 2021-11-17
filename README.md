# parts
Simple inventory management system

## Installation

# MacOS

These steps might get you up and running:

1. Install PostgreSQL by running `brew install postgresql`.
1. Optionally, make PostgreSQL start at boot, with `brew services start postgresql`.
1. Install the migrate CLI with `brew install golang-migrate`.
1. Create the parts database with `createdb parts` (or whatever you want to call the
   database).
1. Run `psql parts -c 'select version();'` to verify that PostgreSQL is up and
   running, and that the database exists.
1. Clone the `parts` repository with `git clone git@github.com:kluzzebass/parts.git`
1. Set up the required environment variables, either in the shell, or by creating
   a `.env` file in the cloned repository. The only variable needed right now is `DB_URL=postgresql://localhost/parts?sslmode=disable`.
1. Run `./migrate.sh up` to bootstrap an empty database.
1. Run `./run.sh` to start the server and hit `http://localhost:8080/` with a browser
   to activate the GraphQL Playground.
1. If any of these steps fail, view it as an exercise in trouble shooting and figure
   it out by yourself. What am I, your mother?

# Windows

I don't know what to do here, but it's probably really difficult and/or annoying.

# Linux

If you're using Linux, I bet you're 31337 enough to figure it out by looking at
the MacOS instructions.

