WEB_DIR=/srv/http


go install
cp config.ini ../../../../bin/config.ini
cp www $WEB_DIR -R
ln manifest $WEB_DIR/www/manifest.json
