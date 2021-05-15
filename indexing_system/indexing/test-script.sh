#!/bin/bash
#$ -M fgomulka@nd.edu
#$ -q *@@jung
#$ -m abe
#$ -r y
#$ -N IndexingTest

make clean
make

## STEP 1: Test with different values of K, using the small files example

# TODO python script that prints out average + stdev times of small files. 
# Steps:
# loop
#   1. turn whatever file into 1 line of records in example.csv
#   2. run it 500 times on k=1, 2, 3, 5, 10, 20
#   3. store these values in a hash table of {k: [500 vals from first iteration + 500 vals from second +...]} etc. for each file.]
# end loop
# take averages/stdev for each k
# print all the data out


## STEP 2: Test with different numbers of inserted records. e.g., when K=3, every time you load 200 records from the example.csv, you can pause and run Top-K algorithm to see how long it takes to find 3 most similar records; 

# TODO I am confused here.... How is STEP 2 different from STEP 3?!

## STEP 3: test with different numbers of MinHash values. The examples I shared on the Drive only contain 1000 MinHash values (and 200 MinHash values for small size files). You can load in different numbers of MinHash values with the examples I shared with you. e.g., you can load the first 400 columns, 600 columns, 800 columns, etc.  

# TODO python script that uses the 1000 minhash value file  1, 5, 10, 20

# loop 
#   1. turn whatever data_x_file into 1 line record in example.csv
#   2. fill it with 200, then run it with k = 1, 5, 10, 20
#   3. fill it with 400, then run it with k = 1, 5, 10, 20
#   4. fill it with 600, then run it with k = 1, 5, 10, 20
#   5. fill it with 800, then run it with k = 1, 5, 10, 20
#   6. fill it with 1000, then run it with k = 1, 5, 10, 20
#   (these all have hash table of (filled: {k: [500 vals + 500 vals ...]})
#   7. do loop again for next data_x_file
# end loop
# take averages/stdev for each k
# print all the data out. 

## STEP 4: Still very confused on the outline here. Compare the times of different file sizes? can't we just use data from the other steps here??

# TODO 

