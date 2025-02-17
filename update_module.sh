#!/bin/sh

if [ "$#" -ne 1 ]; then
	echo "Usage: $0 <new module name>"
	exit 1
fi

new_module=$1

find . -type d -name .git -prune -o -type f -exec sed -i '' -e "s|github.com/siobh9/first-htmx|${new_module}|g" {} \;
