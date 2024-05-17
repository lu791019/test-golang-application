cd /app/application
killall 15kw_modbus_mqtt -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./15kw_modbus_mqtt >> ./log/15kw_modbus/$filename.log 2>&1&