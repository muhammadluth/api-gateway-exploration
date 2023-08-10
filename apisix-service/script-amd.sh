docker-compose -p docker-apisix up -d

args=$(printf "%s" "$1" | tr '[:upper:]' '[:lower:]')
if [ $args=="up" ]
then
    docker-compose -p docker-apisix up -d
else
    docker-compose -p docker-apisix down
fi