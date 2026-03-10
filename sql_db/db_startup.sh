docker run \
    --name mydbserver \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=secret \
    -d postgres 2>/dev/null
psql -h localhost -U postgres -p 5432