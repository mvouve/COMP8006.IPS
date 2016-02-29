WEB_DIR=/srv/http/www


go install
cp config.ini ../../../../bin/config.ini
cp web $WEB_DIR -R
