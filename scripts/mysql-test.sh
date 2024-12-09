#!/bin/bash
mysql -h 127.0.0.1 -P 23306 -u root -D devtool --select-limit=200 --safe-updates --password="test"
