#! /bin/fish


set CONTAINERS (docker ps -a -q)
if set -q "$CONTAINERS"; echo "no containers to remove" ; else ; echo "removing containers"  && docker rm -f "$CONTAINERS" ; end

set IMAGES (docker images -q)
if set -q "$IMAGES"; echo "no images to remove" ; else ; echo "removing images"  && docker rmi -f "$IMAGES" ; end
set VOLUMES (docker volume ls -q)

if set -q "$VOLUMES"; echo "no volumes to remove" ;
else
    echo "removing volumes"
    for volume in $VOLUMES
        docker volume rm $volume
    end
end

