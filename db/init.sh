#bin/bash

cd "$(dirname "$0")"
mysql -h 127.0.0.1 -uroot -p test_db < init.sql