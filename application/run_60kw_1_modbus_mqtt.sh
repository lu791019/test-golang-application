cd /app/application
killall 60kw_1_modbus_mqtt -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./60kw_1_modbus_mqtt >> ./log/60kw_1_modbus/$filename.log 2>&1&