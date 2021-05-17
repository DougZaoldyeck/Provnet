#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N Tests_Task_1

make

## STEP 1: Test with different values of K, using the small files example

# task1.py is a python script that prints out average + stdev times of small files. 
# Steps:
# loop
#   1. turn whatever file into 1 line of records in example.csv
#   2. run it 500 times on k=1, 2, 3, 5, 10, 20
#   3. store these values in a hash table (dict) of {k: [500 vals from first iteration + 500 vals from second +...]} etc. for each file.]
# end loop
# take averages/stdev for each k
# print all the data out

./task1.py > output/task1.txt
