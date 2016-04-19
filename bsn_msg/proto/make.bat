 
 
mkdir %GOPATH%\src\bsn_msg_common
protoc --go_out=%GOPATH%\src  .\bsn_msg_common\*.proto


 
mkdir %GOPATH%\src\bsn_msg_gate2server
protoc --go_out=%GOPATH%\src  .\bsn_msg_gate2server\*.proto
 