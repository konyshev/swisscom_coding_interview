(mkdir -p Docker/tmp_input && cp data/export_SwisscomDB.json swisscom_comodir_restapi config.toml Docker/tmp_input)
(cd Docker && docker-compose up --build --detach)
docker exec container_mongodb mongoimport --db SwisscomDB --collection instances --file /home/input/export_SwisscomDB.json
docker exec container_mongodb ./swisscom_comodir_restapi
rm -rf Docker/tmp_input
