cd /$GOPATH/bin
nohup $GOPATH/bin/router -c $GOPATH/bin/router.conf 2>&1 > $GOPATH/bin/panic-router.log &
nohup $GOPATH/bin/logic -c $GOPATH/bin/logic.conf 2>&1 > $GOPATH/bin/panic-logic.log &
nohup $GOPATH/bin/comet -c $GOPATH/bin/comet.conf 2>&1 > $GOPATH/bin/panic-comet.log &
nohup $GOPATH/bin/job -c $GOPATH/bin/job.conf 2>&1 > $GOPATH/bin/panic-job.log &




{"ver":1,"op":7,"seq":1,"body":{"roomid":37611,"uid":12345}}

{"ver":1,"op":2,"seq":2,"body":{}}

{"ver":1,"op":19,"seq":2,"body":{}}


{"ver":1,"op":15,"seq":1,"body":{"roomid":37611,"uid":12345}}