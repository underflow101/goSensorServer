#! /bin/bash

function start_db() {
    echo "Start Database..."
    echo "Initializing MongoDB..."

    mongod --dbpath /home/pi/db

    echo "Done."
}

case "$1" in
	start)
		start_db
		;;
	stop)
		killall mongo
		echo "Stop DB service."
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