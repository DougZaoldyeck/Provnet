#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N Test_Task_4

make clean
make

## STEP 4: for each file category, run 5 different files with k=1,2,3 as follows

# TODO python script

# loop for each file type:
#   loop for each of first 5 files in the file type
#       1. turn data_x_0 file into 1 line record
#       2. fill system with all mh_values, then run it with k = 1,2,3, 500 times each
#       3. (hash table (dictionary) of values)
#       
#   endloop 
#   take averages + stdevs
# end loop
# compare file types
# 

./task4.py > output/task4.txt
