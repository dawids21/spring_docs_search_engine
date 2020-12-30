#!/bin/sh

CLASSES_DIR=~/Files/Programming/My_programs/spring_docs_scraper/classes_urls

PROJECT=$(ls $CLASSES_DIR | awk 'BEGIN { FS = "." } ; { print $1 }' | rofi -dmenu -p "Choose project: " -i -no-custom)

[[ -z $PROJECT ]] && exit

KEY=$(awk 'BEGIN { FS = "=" } ; { print $1 }' $CLASSES_DIR/$PROJECT.txt | rofi -dmenu -p "Choose class: " -i -no-custom)
LINK=$(grep -oP "^$KEY=\K.+" $CLASSES_DIR/$PROJECT.txt)

[[ -z $LINK ]] && exit

$BROWSER $LINK
