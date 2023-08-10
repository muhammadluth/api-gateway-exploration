args=$(printf "%s" "$1" | tr '[:upper:]' '[:lower:]')
if [ $args=="up" ]
then
    docker-compose -p docker-apisix -f docker-compose-arm64.yml up -d
else
    docker-compose -p docker-apisix -f docker-compose-arm64.yml down
fi