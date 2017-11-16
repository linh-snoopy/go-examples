#!/bin/bash
# a sample for testing goagen cmd
# Get input package design
read -p "Enter your design package flag: " design_path
if [ "$design_path" == "" ]; then
	design_path="github.com/linh-snoopy/go-examples/goatest/design"
fi
echo "Running generator: $design_path"
# Create gen and controller folders
declare -a folders=("controllers" "gen")
for f in "${folders[@]}"
do
	if [ ! -d "$f" ]; then
		echo "Create folder: $f"
		mkdir -p "$f"
	fi
done
# For app, client, swagger
echo "Generate app, client and swagger"
declare -a arr=("app" "client" "swagger")
for sub in "${arr[@]}"
do
	echo "----------- $sub -----------"
	./goagen.exe "$sub" -d "$design_path" -o gen
done
# For controller
if [ -e "main.go" ]; then
	echo "Regenerate controllers"
	./goagen.exe controller -d "$design_path" -o controllers --regen
	for f in "controllers"/*
	do
		sed -i -e '\|"github.com/linh-snoopy/go-examples/goatest/controllers/app"|d' "$f"
	done
else 
	echo "Generate main and controllers"
	./goagen.exe main -d "$design_path" -o controllers
	# copy file main.go out of controller package
	cp controllers/main.go .
	rm controllers/main.go
	sed -i -e 's#/controllers/app#/gen/app#g' main.go
	sed -i '9i\\t"github.com/linh-snoopy/go-examples/goatest/controllers"' main.go
	for f in "controllers"/*
	do
		sed -i -e 's#/controllers/app#/gen/app#g' "$f"
		sed -i -e 's/package main/package controllers/g' "$f"
	done
fi