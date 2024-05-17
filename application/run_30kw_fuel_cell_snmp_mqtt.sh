cd /app/application
killall 30kw_fuel_cell_snmp_mqtt -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./30kw_fuel_cell_snmp_mqtt >> ./log/30kw_fuelcell/$filename.log 2>&1&