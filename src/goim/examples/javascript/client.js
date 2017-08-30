(function(win) {
    var Client = function(options) {
        var MAX_CONNECT_TIME = 10;
        var DELAY = 15000;
        this.options = options || {};
        this.createConnect(MAX_CONNECT_TIME, DELAY);
    }

    Client.prototype.createConnect = function(max, delay) {
        var self = this;
        if (max === 0) {
            return;
        }
        connect();

        var heartbeatInterval;

        function connect() {
            var ws = new WebSocket('ws://123.125.29.128:7272/sub');
            // var ws = new WebSocket('ws://127.0.0.1:8090/sub');
            var auth = false;

            ws.onopen = function() {
                getAuth();
            }

            ws.onmessage = function(evt) {
                var receives = JSON.parse(evt.data)
                for(var i=0; i<receives.length; i++) {
                    var data = receives[i]
                    if (data.op == 8) {
                        auth = true;
                        heartbeat();
                        heartbeatInterval = setInterval(heartbeat, 4 * 60 * 1000);
                    }
                    if (!auth) {
                        setTimeout(getAuth, delay);
                    }
                    if (auth && data.op == 5) {
                        var notify = self.options.notify;
                        if(notify) notify(data.body);
                    }
                }
            }

            ws.onclose = function() {
                if (heartbeatInterval) clearInterval(heartbeatInterval);
                setTimeout(reConnect, delay);
            }

            function heartbeat() {
                ws.send(JSON.stringify({
                    'ver': 1,
                    'op': 2,
                    'seq': 2,
                    'body': {}
                }));
            }

            function getAuth() {
                ws.send(JSON.stringify({
                    'ver': 1,
                    'op': 7,
                    'seq': 1,
                    'body': {
                        'roomid':37611,                        
                        'uid':12345
                    }
                }));
            }

        }

        function reConnect() {
            self.createConnect(--max, delay * 2);
        }
    }

    win['MyClient'] = Client;
})(window);
