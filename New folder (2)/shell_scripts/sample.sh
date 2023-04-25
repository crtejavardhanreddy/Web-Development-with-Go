#!/usr/bin/bash
echo "welcome to the first programming"
echo "the first program to check greatest among two numbers"
read -p "Enter the number x: " x
read -p  "Enter the number y: " y
if [ $x -gt $y ]
then
    echo "x is greater than y"
else
    echo "y is greater than x"
fi