#!/bin/bash

for file in *.json; do
	#echo $file
	new=`echo ${file%.*}_ex.json`
	#echo $new
	cat "$file" | jq . | unexpand -t 2 > "$new"
done

