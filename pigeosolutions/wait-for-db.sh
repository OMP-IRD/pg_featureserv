#!/bin/sh
# Wait for the DB to accept connection from the tileserv role
# Uses the connection string, provided as env var for tileserv
IS_READY=1

while  ! [[ $IS_READY == 0 ]]
do
    echo "Waiting for database to be ready..."
    sleep 5
    psql $DATABASE_URL -c "select version();" > /dev/null 2>&1
#    psql $DATABASE_URL -c "select version();" > /dev/null
    IS_READY=$?
done

# Run pg_featureserv
/pg_featureserv
