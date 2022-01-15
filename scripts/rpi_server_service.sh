#! /bin/bash

function start_agent() {
    echo "[SCRIPT] Start Agent..."
    echo "[SCRIPT] Initializing Eagle Agent..."

    sudo /home/pi/eagle_agent/eagle-agent

    echo "[SCRIPT] Done."
}

case "$1" in
	start)
		start_agent
		;;
	stop)
		killall agent
		killall server
		echo "Stop Agent."
		;;
	restart|reload)
		$0 stop
		$0 start
		;;
	*)
		echo "Usage: $0 {start|stop|restart}"
		exit 1
esac
exit 0