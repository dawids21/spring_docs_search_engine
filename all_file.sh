#!/bin/sh
[[ -f classes_urls/All.txt ]] && rm classes_urls/All.txt
cat classes_urls/*.txt > classes_urls/All.txt
sort -o classes_urls/All.txt classes_urls/All.txt
