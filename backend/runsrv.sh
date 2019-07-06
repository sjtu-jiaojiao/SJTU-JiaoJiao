consul agent -ui -bind=127.0.0.1 -dev > /dev/null 2>&1 &
sleep 5
micro --registry=consul api --handler=http > /dev/null 2>&1 &
sleep 5
micro --registry=consul web > /dev/null 2>&1 &
go run servedoc.go --registry=consul > /dev/null 2>&1 &
consul kv import @consul.json > /dev/null 2>&1 &
goconvey -port 8400 -launchBrowser=false > /dev/null 2>&1 &
realize start > /dev/null 2>&1 &