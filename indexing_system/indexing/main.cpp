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
#include "topK.h"

/* Usage: ./index <example.csv> <number of hash functions (i.e. k)> <new_record.csv> <topk num> */

using std::ifstream;
using std::stringstream;
using std::getline;
using std::istringstream;
using std::cerr;


int main(int argc, char *argv[]) {
	//int k = 800; // Number of hash functions.
	string filename; 

	// handle arguments
	if (argc != 5) {
		cout << "wrong number of args. see usage." << endl;
		return 1;
	}
	filename = argv[1];

	istringstream ss(argv[2]);
	int k; // number of hash functions
	if (!(ss >> k)) {
		cerr << "Invalid number: " << argv[2] << '\n';
	} else if (!ss.eof()) {
		cerr << "Trailing characters after number: " << argv[2] << '\n';
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
			}
			hf++;
		}
    }

    // convert all fll to vectors
    system.convert();


	// Find similarities across *new* data
	// First, process new data from a new file
	string newfilename = argv[3];
	
	// handle csv file 
	ifstream nfs;
	nfs.open(newfilename);

	// error message if file does not exist in directory
    if (!nfs) {
        cout << "error opening new file " << newfilename << endl;
        return 1;
    }
	
	string nline, nrecord, nmh_val;

	// read in csv file to indexing system. also find list record ids with
	// same minhash value!
	vector<string> similar; // vector of all similar records
    while (!nfs.eof()) { // TODO this should only be one iteration. What else to check?
		getline(nfs, nline);
		stringstream s(nline);
		int hf = 0; // hash function number

		//read each column
		getline(s, nrecord, ',');
		cout << "new record: " << nrecord << endl;
		while (getline(s, nmh_val, ',')){
			if (nmh_val != ""){
				system.add_record(hf, nrecord, nmh_val); // TODO why are we still using linked list? kinda useless at this point, correct?
				system.next_record(hf, nrecord, nmh_val, similar); // add record to respective vector

			}
			hf++;
		}
    }

    // now that we have the vector of vectors of similar records, get the top K similar! 
    vector<string> results; // initiate the final array of most similar records


	// The code below is for some vector 
	istringstream ss_k(argv[4]);
	int tk;
	if (!(ss_k >> tk)) {
		cerr << "Invalid number: " << argv[4] << '\n';
	} else if (!ss_k.eof()) {
		cerr << "Trailing characters after number: " << argv[4] << '\n';
	}

	TopKElements top_k;
	results = top_k.topKFrequent(similar, tk);

	for (int i = 0; i < results.size(); i++) {
		if (i == results.size() - 1)
			cout << results[i] << endl;
		else
			cout << results[i] << ",";
	}




	return 0;
}
