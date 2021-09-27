#!/usr/bin/env bash
wc -w heidi.txt | awk '{print $1}' | tr -d '\n'
