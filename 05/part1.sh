#!/bin/bash
cat i.txt | tr "FLBR" "0011" | while read nn ;do echo "$((2#$nn))";done | sort -n | tail -n 1


