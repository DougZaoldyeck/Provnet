#!/usr/bin/env python3

from subprocess import Popen, PIPE
import numpy
import sys

## GLOBALS ##

K_VALS = [1, 5, 10, 20]

RECORDS = [200, 400, 600, 800, 1000]

K_INSERT_DICT =    {   
                200: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                400: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                600: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                800: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                1000: {1: [],
                    5: [],
                    10:[],
                    20:[],},
            }

K_FINDTK_DICT =    {   
                200: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                400: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                600: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                800: {1: [],
                    5: [],
                    10:[],
                    20:[],},
                1000: {1: [],
                    5: [],
                    10:[],
                    20:[],},
            }

## FUNCTIONS ##

def run_system(i):
    for j in K_VALS: # for each k
        for records in RECORDS:
            for _ in range(500): # every k, 500 times each
                process = Popen([f'./main', f'../testing/10-100mb/combine.csv', f'1000', f'../testing/10-100mb/input/data_{i}_0.csv', f'{j}', f'{records}'], stdout=PIPE, stderr=PIPE)
                (output, err) = process.communicate()
                exit_code = process.wait()
                #output = output.rstrip().split()
                #print(err)
                try:
                    K_INSERT_DICT[records][j].append(int(output.rstrip().split()[0]))
                    K_FINDTK_DICT[records][j].append(int(output.rstrip().split()[1]))
                except IndexError:
                    print(f'output: {output}')
                    print(f'file no: {i}, k-val: {j} WTH')
                    print(err)
                    print(output)

## MAIN FUNCTION ##

def main():
    
    for i in range(21, 31): # for each file
        run_system(i)

    for records in RECORDS:
        print(f'~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
        print(f'NUM RECORDS:              {records}')
        print(f'~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
        for k in K_VALS:
            print(f'K-VALUE:                {k}')
            print(f'--------INSERTION-----------')
            print(f'AVERAGE (microseconds): {numpy.average(K_INSERT_DICT[records][k])}')
            print(f'STD DEV (microseconds): {numpy.std(K_INSERT_DICT[records][k])}')
            print(f'--------FINDING TOPK--------')
            print(f'AVERAGE (microseconds): {numpy.average(K_FINDTK_DICT[records][k])}')
            print(f'STD DEV (microseconds): {numpy.std(K_FINDTK_DICT[records][k])}')
            print()

    original_stdout = sys.stdout 
    with open('./output/task2/insertion_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_INSERT_DICT}')
        sys.stdout = original_stdout
    
    with open('./output/task2/findtopk_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_FINDTK_DICT}')
        sys.stdout = original_stdout

## MAIN ##

if __name__ == "__main__":
    main()
