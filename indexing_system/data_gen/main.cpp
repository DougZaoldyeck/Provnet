#include <iostream>
#include <fstream>
#include <string>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <math.h>
#include <sstream>
#include <vector>
#include <string.h>
#include <cstdlib>
#include "sha256.h"
#include <time.h>

using std::cout;
using std::endl;

using namespace std;

int main(int argc, char *argv[]) {
    clock_t t;
    t=clock();
    //char filename[100];
    //filename = argv;
    std::ofstream myfile;
    string myname = argv[1];
    string csv = ".csv";
    myname = myname + csv;
    myfile.open (myname, std::ios_base::app);

    FILE *f;
    f = fopen(argv[1], "r");
    if (f == NULL){
        printf("Error! opening file\n");
    exit(1);
    }

    char str[1000];
    string min_hash[200];
    for (int y=0; y<200; y++)
        min_hash[y]="ffff";
    string output;


    while (fgets(str, 1000, f)){        //means we only calculate the "data" segment as each element
        if (str[3]=='d' && str[4]=='a' && str[5]=='t' && str[6]=='a'){
            for (int n=0; n<200; n++){
                output = sha256(str+n);
                if (output < min_hash[n])
                    min_hash[n] = output;
            }
            while (fgets(str, 1000, f)){
                int len = strlen(str);
                if (len >3){
                    for (int n=0; n<200; n++){
                        output = sha256(str+n);
                        if (output < min_hash[n])
                            min_hash[n] = output;
                    }
                    continue;
                }
                else break;             //means we reach the end of json file
            }
        }
        continue;
    }



    for (int y = 0; y < 200; y++)
    {
        myfile << min_hash[y] << ",";
    }
    myfile << '\n';
    myfile.close();


    fclose(f);

    t = clock() - t;
    double time_taken = ((double)t)/CLOCKS_PER_SEC; // calculate the elapsed time
    printf("The program took %f seconds to execute\n", time_taken);

    return (0);
}


