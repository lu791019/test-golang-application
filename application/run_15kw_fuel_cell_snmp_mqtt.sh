cd /app/application
killall 15kw_fuel_cell_snmp_mqtt -q
filename=$(date +%Y%m%d)
#go build
sleep 1
nohup ./15kw_fuel_cell_snmp_mqtt >> ./log/15kw_fuelcell/$filename.log 2>&1&