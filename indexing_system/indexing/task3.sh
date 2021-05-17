#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N Test_Task_3

make clean
make

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
