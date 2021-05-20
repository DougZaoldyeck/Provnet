#!/usr/bin/env python3

from subprocess import Popen, PIPE
import numpy
import sys

## GLOBALS ##

K_VALS = [1, 2, 3]
SIZES = ['sml', 'med', 'lrg']
K_INSERT_DICT =    {    'sml': {
                            1: [],
                            2: [],
                            3: [],},
                        'med': {
                            1: [],
                            2: [],
                            3: [],},
                        'lrg': {
                            1: [],
                            2: [],
                            3: [],},
                    }

K_FINDTK_DICT =    {    'sml': {
                            1: [],
                            2: [],
                            3: [],},
                        'med': {
                            1: [],
                            2: [],
                            3: [],},
                        'lrg': {
                            1: [],
                            2: [],
                            3: [],},
                    }
## FUNCTIONS ##

def run_sml(i):
    for j in K_VALS: # for each k
        for _ in range(500): # every k, 500 times each
            process = Popen([f'./main', f'../testing/40-200kb/Small-file-examples.csv', f'1000', f'../testing/40-200kb/input/data_{i}_0.csv', f'{j}', f'1000'], stdout=PIPE)
            (output, err) = process.communicate()
            exit_code = process.wait()
            output = output.rstrip().split()
            K_INSERT_DICT['sml'][j].append(int(output[0]))
            K_FINDTK_DICT['sml'][j].append(int(output[1]))


def run_med(i):
    for j in K_VALS: # for each k
        for _ in range(500): # every k, 500 times each
            process = Popen([f'./main', f'../testing/10-100mb/Middle-file-examples.csv', f'1000', f'../testing/10-100mb/input/data_{i}_0.csv', f'{j}', f'1000'], stdout=PIPE)
            (output, err) = process.communicate()
            exit_code = process.wait()
            output = output.rstrip().split()
            K_INSERT_DICT['med'][j].append(int(output[0]))
            K_FINDTK_DICT['med'][j].append(int(output[1]))


def run_lrg(i):
    for j in K_VALS: # for each k
        for _ in range(500): # every k, 500 times each
            process = Popen([f'./main', f'../testing/100mb-1gb/large-file-examples.csv', f'1000', f'../testing/100mb-1gb/input/data_{i}_0.csv', f'{j}', f'1000'], stdout=PIPE)
            (output, err) = process.communicate()
            exit_code = process.wait()
            output = output.rstrip().split()
            K_INSERT_DICT['lrg'][j].append(int(output[0]))
            K_FINDTK_DICT['lrg'][j].append(int(output[1]))

## MAIN FUNCTION ##

def main():
    
    for i in range(1, 6): # 5 small files
        run_sml(i)
    for i in range(21, 26): # 5 medium files
        run_med(i)
    for i in range(31, 36): # 5 large files
        run_lrg(i)

    for size in SIZES:
        print(f'~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
        print(f'FILE SIZE:              {size}')
        print(f'~~~~~~~~~~~~~~~~~~~~~~~~~~~~')
        for k in K_VALS:
            print(f'K-VALUE:                {k}')
            print(f'--------INSERTION-----------')
            print(f'AVERAGE (microseconds): {numpy.average(K_INSERT_DICT[size][k])}')
            print(f'STD DEV (microseconds): {numpy.std(K_INSERT_DICT[size][k])}')
            print(f'--------FINDING TOPK--------')
            print(f'AVERAGE (microseconds): {numpy.average(K_FINDTK_DICT[size][k])}')
            print(f'STD DEV (microseconds): {numpy.std(K_FINDTK_DICT[size][k])}')
            print()

    original_stdout = sys.stdout 
    with open('./output/task4/insertion_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_INSERT_DICT}')
        sys.stdout = original_stdout
    
    with open('./output/task4/findtopk_data.json', 'w') as f:
        # NOTE each set of 500 data points is from each file.
        sys.stdout = f # Change the standard output to the file we created.
        print(f'{K_FINDTK_DICT}')
        sys.stdout = original_stdout

## MAIN ##

if __name__ == "__main__":
    main()
