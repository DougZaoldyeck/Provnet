#!/usr/bin/env python3


import os

for i in range(21, 31):
    os.system(f'../data_gen/testoutput ../testing/10-100mb/data_{i}_0')
for i in range(31, 41):
    os.system(f'../data_gen/testoutput ../testing/100mb-1gb/data_{i}_0')
