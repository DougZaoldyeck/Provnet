/* This is an implementation of the indexing system, as described by
 * Changhao Chenli 
 * 
 * Frank Gomulka
 */
#include <iostream>
#include <fstream>
#include <stdio.h>
#include <math.h>
#include <map>
#include <sstream>
#include <string>
#include "sha256.h"
#include "indexing.h"

/* Usage: ./index <example.csv> <number of hash functions (i.e. k)> */

using std::ifstream;
using std::stringstream;
using std::getline;


int main(int argc, char *argv[]) {
	int k = 800; // Number of hash functions.
	string filename; 

	// handle arguments
	if (argc < 3 || argc > 3) {
		cout << "no filename selected. see usage." << endl;
		return 1;
	}
	filename = argv[1];

	std::istringstream ss(argv[2]);
	int x;
	if (!(ss >> x)) {
		std::cerr << "Invalid number: " << argv[2] << '\n';
	} else if (!ss.eof()) {
		std::cerr << "Trailing characters after number: " << argv[2] << '\n';
	}
	
	// create indexing system
	Indexing system = Indexing(k);

	// handle csv file 
	ifstream ifs;
	ifs.open(filename);

	// error message if file does not exist in directory
    if (!ifs) {
        cout << "error opening file " << filename << endl;
        return 1;
    }
	
	string line, mh_val, record;


	// read in csv file to indexing system.
    while (!ifs.eof()) {
		getline(ifs, line);
		stringstream s(line);
		int hf = 0; // hash function number

		//read each column
		getline(s, record, ',');
		while (getline(s, mh_val, ',')){
			if (mh_val != ""){
				system.add_record(hf, record, mh_val);
				//system.list_records(hf, mh_val);

				//last_hf = hf;
				//last_mh = mh_val;
			}
			hf++;
		}
    }

	// Find similarities across certain lines of data...

	/*
	// The code below is for some vector 


	*/


	return 0;
}
