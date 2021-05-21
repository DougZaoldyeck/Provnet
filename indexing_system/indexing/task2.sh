#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N Test_Task_2

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
