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
error_handle() {
	if [ $? -ne 0 ]; then
		echo "Failed and stop bash"
		exit 1
	fi
}
# Detect ostype
os_detector() {
	if [[ "$OSTYPE" == "linux-gnu" ]]; then
		cmd="goagen"
	elif [[ "$OSTYPE" == "cygwin" ]]; then
		# POSIX compatibility layer and Linux environment emulation for Windows
		cmd="./goagen.exe"
	elif [[ "$OSTYPE" == "msys" ]]; then
		# Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
		cmd="./goagen.exe"
	elif [[ "$OSTYPE" == "darwin"* ]]; then
		# MAC OSX
		cmd="goagen"
	else
		cmd="goagen"
	fi
}
# For app, client, swagger
os_detector
echo $cmd
echo "Generate app, client and swagger"
declare -a arr=("app" "client" "swagger" "js")
for sub in "${arr[@]}"
do
	echo "----------- $sub -----------"
	"$cmd" "$sub" -d "$design_path" -o gen
	error_handle
done
# For controller
if [ -e "main.go" ]; then
	echo "Regenerate controllers"
	"$cmd" controller -d "$design_path" -o controllers --regen
	error_handle
	for f in "controllers"/*
	do
		sed -i -e '\|"github.com/linh-snoopy/go-examples/goatest/controllers/app"|d' "$f"
	done
else 
	echo "Generate main and controllers"
	"$cmd" main -d "$design_path" -o controllers
	error_handle
	# copy file main.go out of controller package
	mv controllers/main.go .
	sed -i -e 's#/controllers/app#/gen/app#g' main.go
	sed -i '9i\\t"github.com/linh-snoopy/go-examples/goatest/controllers"' main.go
	for f in "controllers"/*
	do
		sed -i -e 's#/controllers/app#/gen/app#g' "$f"
		sed -i -e 's/package main/package controllers/g' "$f"
	done
fi