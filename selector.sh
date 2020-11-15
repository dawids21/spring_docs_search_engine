#!/bin/sh

PROJECT=$(ls classes_urls | awk 'BEGIN { FS = "." } ; { print $1 }' | rofi -dmenu -p "Choose project: " -i -no-custom)

[[ -z $PROJECT ]] && exit

KEY=$(awk 'BEGIN { FS = "=" } ; { print $1 }' classes_urls/$PROJECT.txt | rofi -dmenu -p "Choose class: " -i -no-custom)
LINK=$(grep -oP "^$KEY=\K.+" classes_urls/$PROJECT.txt)

[[ -z $LINK ]] && exit

qutebrowser $LINK
