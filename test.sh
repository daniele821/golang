#!/bin/bash

echo -n "STDERR TESTING HERE!" >&2
echo -n "fuck you: "
read -r line
echo "$line"
