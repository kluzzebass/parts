#!/bin/bash

function usage {
  echo
  echo "$0 <command> [arguments...]"
  echo
  echo "  Commands:"
  echo
  echo "    create <migration name>"
  echo "    goto <version>"
  echo "    up [<number of migrations>]"
  echo "    down [{<number of migrations>|-all}]"
  echo "    drop [-f]"
  echo "    force <version>"
  echo "    version"
}

MIGRATE="$( command -v migrate )"
MIGRATIONS="db/migrations"

if ! [ -x "$MIGRATE" ]; then
  echo The migrate command is missing! Check https://github.com/golang-migrate/migrate/tree/master/cmd/migrate for details.
  exit 1
fi

if [ -f .env ]; then
  source .env
fi

if [ -z "${DB_URL}" ]; then
  echo Please set the DB_URL environment variable! Check https://github.com/golang-migrate/migrate/tree/master/cmd/migrate for details.
  exit 1
fi

case $1 in

  "create")
    $MIGRATE create -ext sql -dir $MIGRATIONS -seq $2
    ;;

  "goto")
    $MIGRATE -path $MIGRATIONS -database $DB_URL goto $2
    ;;

  "up")
    $MIGRATE -path $MIGRATIONS -database $DB_URL up $2
    ;;

  "down")
    $MIGRATE -path $MIGRATIONS -database $DB_URL down $2
    ;;

  "drop")
    $MIGRATE -path $MIGRATIONS -database $DB_URL drop $2
    ;;

  "force")
    $MIGRATE -path $MIGRATIONS -database $DB_URL force $2
    ;;

  "version")
    $MIGRATE -path $MIGRATIONS -database $DB_URL version
    ;;

  *)
    usage
    exit 1
    ;;
esac






