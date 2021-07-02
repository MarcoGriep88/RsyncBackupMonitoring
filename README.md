# Rsync Backup Monitoring
Scripts and APIs for Centralized Monitoring of Rsync Backups

# IMPORT INFORMATION
I CURRENTLY WORK ON AN EXTENDED VERSION OF THIS SOFTWARE. THERE WILL BE AN NEW VERSION ON GITLAB SOON (I will link it here). THE NEW VERSION WILL NOT BE COMPATIBLE BECAUSE I COMPLETLY REDESIGN THIS APPLICATION. WHAT WILL BE NEW:

* Login
* User Management
* Better Parsing from Logfiles
* Better Docker-Compose Setup
* VueJS Application will be Redesigned and Work with Vuetify
* Calender Integration
* Logfiles will be uploaded so you can delete them on your host

## What does this Software do?
This Software (main.go) will check your Rsync Logfiles, analyse them and Report your Rsync Backup Results to an Web API (web). You can monitor your Backups by the Web Frontend (app-dashboard).

## Info about Sponsoring
This Software lives from Contribution and Support of the Community. If you use this set of Software or find it useful, please consider to donate by the "Sponsoring Button above" (Patreon or PayPal). By Sponsoring this Project you are helping me financing Hosting and Infrastrucuter Costs, help to to improve my skills (Education) and have more freetime to work on this Project. Thanks!

## How to install this Software?
I highly recommend to use the rebuild.sh Script in combination with docker and docker-compose. 

Steps to install:
* Install docker and docker-compose on your linux machine.
* Install go-lang on your linux machine
* Install npm and nodejs
* compile the main.go file with:

```
go main.go
```
* Rename the compiled main.go file "RsyncBackupReport"
* Install Rsync on your linux machine (if not already installed)
* Create an Rsync Backup cronjob on your linux machine (the actual backup)

Example:
crontab -e
```
0 10 * * * /bin/bash /home/user/backup_ipsl001.sh >> /home/user/logs/backup.log 2>&1
```

backup_ipsl001.sh
```
sshpass -p "MYPASSWORD" rsync --stats --log-file /home/user/logs/$(date +%Y-%m-%d)_volumes.log -vazP --delete /var/lib/docker/volumes admin@192.168.112.200:/share/Public/ipsl001

/home/user/RsyncBackupReport /home/user/logs/$(date +%Y-%m-%d)_volumes.log "http://localhost:15000" "Mirror-Volumes"
```
* Replace "http://localhost:15000" with the IP of your Server where the API is running.

After the Rsync Job the RSyncBackupReport will analyse the logfile (Start argument 1) and send the Data to the Web API. "Mirror-Volumes" is the Job Description for the Frontend.

* Open the Directory "App/app-dashboard/public-html"

* Open the vue.config.js file and replace the IP and Port with the IP and Port of your server (Not localhost.) Each Client in you network need to be able to connect to this address.

* Do the same thing in the directory /App/app-dashboard/public-html/src/components/server/BackupOverview.vue
Replace the IP in the JavaScript at the bottom of the code

* In the Directory /App/app-dashboard/public-html run:

```
npm run build
```

* Now run the rebuild.sh Script

Your Installation should work. Check if you can reach localhost:15000 (You should get an empty json response). Also Check localhost:15100 for the Frontend.

If both work. Do an Test backup and check if the job is listed after the finished backup


[German Blog Article about this Rsync Backup Monitoring](https://www.protoncode.eu/post/zentrale-rsync-backup-ueberwachung-mit-go-lang-und-vuejs/)


## Release Notes

### 2020-06-18
* Script now reports each file the monitoring server that has been in backup
* Changed Routing in VueJS Application to vue-router
* Created an Component in VueJS Application to display each backuped filed
* Created an Rebuild Shellscript for the App-Dashboard
* Bugfix in App-Dashboard (VueJS APP) in Table-Ordering

# Contact
If you have any questions, dont hestitate to write me an E-Mail. See the Contact Information on my [IT-Consultant](https://www.marcogriep.de) Webpage.
