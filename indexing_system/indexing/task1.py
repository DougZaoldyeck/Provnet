#!/usr/bin/env python3

from subprocess import Popen, PIPE
import numpy
import sys

## GLOBALS ##

K_VALS = [1, 2, 3, 5, 10, 20]
K_INSERT_DICT =    {   1: [],
                2: [],
                3: [],
                5: [],
                10:[],
                20:[],
            }
K_FINDTK_DICT =    {   1: [],
                2: [],
                3: [],
                5: [],
                10:[],
                20:[],
            }

## FUNCTIONS ##

def run_system(i):
    for j in K_VALS: # for each k
        for _ in range(500): # every k, 500 times each
            process = Popen([f'./main', f'../testing/40-200kb/Small-file-examples.csv', f'1000', f'../testing/40-200kb/input/data_{i}_0.csv', f'{j}', f'1000'], stdout=PIPE)
            (output, err) = process.communicate()
            exit_code = process.wait()
            output = output.rstrip().split()
            try:
                K_INSERT_DICT[j].append(int(output[0]))
            except IndexError:
                print(f'output: {output}')
                print(f'file no: {i}, k-val: {j} WTH')
            K_FINDTK_DICT[j].append(int(output[1]))

## MAIN FUNCTION ##

def main():
    
    #for i in range(1, 21): # for each file
    #    run_system(i)

    for k in K_VALS:
        print(f'K-VALUE:                {k}')
        print(f'--------INSERTION-----------')
        print(f'AVERAGE (microseconds): {numpy.average(K_INSERT_DICT[k])}')
        print(f'STD DEV (microseconds): {numpy.std(K_INSERT_DICT[k])}')
        print(f'--------FINDING TOPK--------')
        print(f'AVERAGE (microseconds): {numpy.average(K_FINDTK_DICT[k])}')
        print(f'STD DEV (microseconds): {numpy.std(K_FINDTK_DICT[k])}')
        print()

    original_stdout = sys.stdout 
    with open('./output/task1/insertion_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_INSERT_DICT}')
        sys.stdout = original_stdout
    
    with open('./output/task1/findtopk_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_FINDTK_DICT}')
        sys.stdout = original_stdout

## MAIN ##

if __name__ == "__main__":
    main()
