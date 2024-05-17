這是一個使用 Golang 程式語言開發的測試專案。

該專案旨在實現與 MQTT(Message Queuing Telemetry Transport) 和 Modbus TCP 協議的通訊。

MQTT 是一種輕量級的發布/訂閱式消息傳遞協議,常用於物聯網(IoT)和機器對機器(M2M)通訊。它可以在低頻寬、不可靠的網路環境下提供高效能的資料傳輸。

Modbus TCP 是基於 Modbus 應用層消息協定的 TCP/IP 版本。它廣泛應用於工業自動化系統中,用於設備監控和控制。透過 Modbus TCP,可以讀取和寫入連線裝置的記憶體區塊,實現對遠端設備的操作和資料獲取。

在這個專案中,我們運用多個與不同電池容量(如 15kW、30kW 和 60kW)相關的檔案和程式碼。
這些檔案包含了與電池管理系統通訊的邏輯,使用 MQTT 協議傳輸電池的狀態和數據,同時使用 Modbus TCP 協議讀取和控制相關的工業設備。

此外,專案還包含了一些其他檔案,如 go.mod(Go 模組定義檔案)、go.sum(Go 模組校驗碼檔案)和 mqtt_mysql.go(可能用於將資料存儲到 MySQL 資料庫)。

Summary ,這個專案展示了如何使用 Go 語言開發一個能夠與 MQTT 和 Modbus TCP 協議互動的應用程式,可應用於工業自動化和物聯網領域,實現設備監控、控制和資料收集等功能。



This is a project developed using the Go programming language. 
The project aims to implement communication with the MQTT (Message Queuing Telemetry Transport) and Modbus TCP protocols. 

MQTT is a lightweight publish/subscribe messaging protocol commonly used in the Internet of Things (IoT) and machine-to-machine (M2M) communications. It can provide efficient data transfer in low-bandwidth and unreliable network environments. 

Modbus TCP is the TCP/IP version of the Modbus application layer messaging protocol. It is widely used in industrial automation systems for device monitoring and control. Through Modbus TCP, it is possible to read and write memory blocks of connected devices, enabling remote device operation and data acquisition.
 
In this project, we utilize multiple files and code related to different battery capacities, such as 15kW, 30kW, and 60kW.

These files contain the logic for communicating with the battery management system, using the MQTT protocol to transmit battery status and data, while also using the Modbus TCP protocol to read and control related industrial equipment.

Additionally, the project includes other files such as go.mod (Go module definition file), go.sum (Go module checksum file), and mqtt_mysql.go (possibly used for storing data in a MySQL database).

In summary, this project demonstrates how to develop an application using the Go language that can interact with the MQTT and Modbus TCP protocols, applicable in industrial automation and the Internet of Things domains, enabling device monitoring, control, and data collection functionalities.
