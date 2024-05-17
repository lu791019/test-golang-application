cd /app/application
killall mqtt_mysql -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./mqtt_mysql >> ./log/mysql/$filename.log 2>&1&