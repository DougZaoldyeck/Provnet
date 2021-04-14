/* This is an implementation of the indexing system, as described by
 * Changhao Chenli 
 * 
 * Frank Gomulka
 */
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include "sha256.h"
#include "indexing.h"

/* Usage: ./index <example.csv> */
using std::ifstream;
using std::stringstream;
using std::getline;


int main(int argc, char *argv[]) {
	int k = 800; // Number of hash functions.
	string filename; 

	// handle arguments
	if (argc < 2 || argc > 2) {
		cout << "no filename selected. see usage." << endl;
		return 1;
	}
	filename = argv[1];
	
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
	// TODO temp
	int last;

    while (!ifs.eof()) {
		getline(ifs, line);
		stringstream s(line);
		int hf = 0; // hash function number

		//read each column
		getline(s, record, ',');
		while (getline(s, mh_val, ',')){
			if (mh_val != ""){
				system.add_record(hf, record, mh_val);
				system.list_records(hf, mh_val);

				last = hf-1;
			}
			hf++;
		}
    }
	

	// TODO testing
	cout << "current structure:" << endl;
	system.list_records(last, mh_val);

	return 0;
}
