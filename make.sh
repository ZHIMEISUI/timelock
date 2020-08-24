go build ./TendermintApp/ABCIServer/
go build ./TendermintApp/ABCIClient/
go build  -o Client ./ClientApp/
go build  -o AliceClient ./AliceClientApp/
go build  -o BobClient ./BobClientApp/
rm -rf log