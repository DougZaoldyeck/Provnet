#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N FinalSemesterTests

make clean
make

## STEP 2: Test with different numbers of inserted records. e.g., when K=3, every time you load 200 records from the example.csv, you can pause and run Top-K algorithm to see how long it takes to find 3 most similar records; 

# TODO python script that uses 1000 minhash value file for 1 5 10 20

# loop 
#   1. turn whatever data_x_file into 1 line record in example.csv
#   2. fill it with 200 records, then run it with k = 1, 5, 10, 20
#   3. fill it with 400 records, then run it with k = 1, 5, 10, 20
#   4. fill it with 600 records, then run it with k = 1, 5, 10, 20
#   5. fill it with 800 records, then run it with k = 1, 5, 10, 20
#   6. fill it with 1000 records, then run it with k = 1, 5, 10, 20
#   (these all have hash table of (filled: {k: [500 vals + 500 vals ...]})
#   7. do loop again for next data_x_file
# end loop
# take averages/stdev for each k
# print all the data out. 

./task2.py > output/task2.txt


## STEP 3: test with different numbers of MinHash values. The examples I shared on the Drive only contain 1000 MinHash values (and 200 MinHash values for small size files). You can load in different numbers of MinHash values with the examples I shared with you. e.g., you can load the first 400 columns, 600 columns, 800 columns, etc.  

# TODO python script that uses the 1000 minhash value file  1, 5, 10, 20

# loop 
#   1. turn whatever data_x_file into 1 line record in example.csv
#   2. fill it with 200 mh_vals, then run it with k = 1, 5, 10, 20, 500 times each
#   3. fill it with 400 mh_vals, then run it with k = 1, 5, 10, 20
#   4. fill it with 600 mh_vals, then run it with k = 1, 5, 10, 20
#   5. fill it with 800 mh_vals, then run it with k = 1, 5, 10, 20
#   6. fill it with 1000 mh_vals, then run it with k = 1, 5, 10, 20
#   (these all have hash table (dict) of (filled: {k: [500 vals + 500 vals ...]})
#   7. do loop again for next data_x_file
# end loop
# take averages/stdev for each k
# print all the data out. 

./task3.py > output/task3.txt

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
