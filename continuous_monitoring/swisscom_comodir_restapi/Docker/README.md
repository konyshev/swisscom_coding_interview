for backup: 
`mongodump -d SwisscomDB -o /data/db_backup/`

for restore: 
`mongorestore -d SwisscomDB /data/db_backup/SwisscomDB/`

----
for export:
`mongoexport --db SwisscomDB --collection instances --out /data/db_backup/export_SwisscomDB.json`

for import:
`mongoimport --db SwisscomDB --collection instances --file /data/db_backup/export_SwisscomDB.json`