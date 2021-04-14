#include <iostream>
#include <fstream>
#include <string>
#include "sha256.h"
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <math.h>
#include <sstream>
#include <vector>

using std::string;
using std::cout;
using std::endl;


//This is for hash calculation
int main()
{
    std::ofstream myfile;
    myfile.open ("example.csv");
    int i; // i represents the number of total records to be inserted
    for(i=0; i<1000; i++)
    {
        std::string str = std::to_string(i);
        string output1 = sha256(str);
        myfile << output1 << ",";
        for(int j=0; j<800; j++) // j means the number MinHash functions our system has
        {
            output1 = output1 + "1";
            string output2 = sha256(output1);
            myfile << output2 << ",";
        }
        myfile << '\n';
    }
    myfile.close();
    return 0;
}
