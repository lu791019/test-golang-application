cd /app/application
killall 30kw_modbus_mqtt -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./30kw_modbus_mqtt >> ./log/30kw_modbus/$filename.log 2>&1&