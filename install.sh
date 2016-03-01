WEB_DIR=/var/www


go build
#cp config.ini ../../../../bin/config.ini
cp html $WEB_DIR -R
ln manifest $WEB_DIR/html/manifest.json
