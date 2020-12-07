#!/bin/bash
old=-1
cat i.txt | tr "FLBR" "0011" | while read nn ;do echo "$((2#$nn))";done | sort -n | while read sn; do d=$((sn-old)); if test $d -eq 2; then echo $((sn-1)); break;fi;old=$sn; done
