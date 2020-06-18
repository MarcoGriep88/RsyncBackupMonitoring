docker stop app-backup-rsync-apache-001
docker rm app-backup-rsync-apache-001
npm run build
docker build -t app-backup-rsync-apache-001 .
docker run -dit --name app-backup-rsync-apache-001 -p 80:80 app-backup-rsync-apache