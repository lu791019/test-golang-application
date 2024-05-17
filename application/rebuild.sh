killall mqtt_mysql -q
killall 15kw_modbus_mqtt -q
killall 30kw_modbus_mqtt -q
killall 60kw_1_modbus_mqtt -q
killall 60kw_2_modbus_mqtt -q
killall 15kw_fuel_cell_snmp_mqtt -q
killall 30kw_fuel_cell_snmp_mqtt -q
rm ./log/*/*.log
go build mqtt_mysql.go
go build 15kw_modbus_mqtt.go
go build 30kw_modbus_mqtt.go
go build 60kw_1_modbus_mqtt.go
go build 60kw_2_modbus_mqtt.go
go build 15kw_fuel_cell_snmp_mqtt.go
go build 30kw_fuel_cell_snmp_mqtt.go
./run_mqtt_mysql.sh
./run_15kw_modbus_mqtt.sh
./run_30kw_modbus_mqtt.sh
./run_60kw_1_modbus_mqtt.sh
./run_60kw_2_modbus_mqtt.sh
./run_15kw_fuel_cell_snmp_mqtt.sh
./run_30kw_fuel_cell_snmp_mqtt.sh