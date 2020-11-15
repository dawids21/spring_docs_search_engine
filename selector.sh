#!/bin/sh
KEY=$(awk 'BEGIN { FS = "=" } ; { print $1 }' scraper/file.txt | rofi -dmenu -p "Choose class: " -i -no-custom)
LINK=$(grep -oP "^$KEY=\K.+" scraper/file.txt)

[[ ! -z $LINK ]] && qutebrowser $LINK
