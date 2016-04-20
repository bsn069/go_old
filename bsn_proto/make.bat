 
rmdir %GOPATH%\src\bsn_common /s /q
mkdir %GOPATH%\src\bsn_common
protoc --go_out=%GOPATH%\src  .\bsn_common\*.proto
 
rmdir %GOPATH%\src\bsn_msg_common /s /q
mkdir %GOPATH%\src\bsn_msg_common
protoc --go_out=%GOPATH%\src  .\bsn_msg_common\*.proto
 
rmdir %GOPATH%\src\bsn_msg_gate_server /s /q
mkdir %GOPATH%\src\bsn_msg_gate_server
protoc --go_out=%GOPATH%\src  .\bsn_msg_gate_server\*.proto

rmdir %GOPATH%\src\bsn_msg_gate_gateconfig /s /q
mkdir %GOPATH%\src\bsn_msg_gate_gateconfig
protoc --go_out=%GOPATH%\src  .\bsn_msg_gate_gateconfig\*.proto
 