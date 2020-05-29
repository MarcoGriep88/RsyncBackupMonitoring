docker-compose down
#git pull 
docker-compose up --build -d
sleep 30
docker exec -i api-backup-rsync-lumen-001 sh -c 'cd /var/www && php artisan migrate'
echo "finished"
sleep 10