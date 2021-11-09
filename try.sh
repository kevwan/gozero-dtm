(cd trans && go run trans.go &)
(cd dtmsvr && go run dtmsvr.go &)
(cd dtmsample && sleep 3 && go run dtmsample.go)
# (cd order && go run order.go &)
# (cd client && go run client.go)
pkill order trans dtmsvr dtmsample
