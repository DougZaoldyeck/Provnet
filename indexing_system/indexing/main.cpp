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
#include <time.h>

#include "sha256.h"
#include "indexing.h"
#include "topK.h"

/* Usage: ./main <current_system.csv> <number of hash functions (i.e. hfn_cap)> <new_record.csv> <topk num> <mh_vak_num> */

using std::ifstream;
using std::stringstream;
using std::getline;
using std::istringstream;
using std::cerr;


int main(int argc, char *argv[]) {
	string filename; 

	// handle arguments
	if (argc != 6) {
		cout << "wrong number of args. see usage." << endl;
		return 1;
	}
	filename = argv[1];

	istringstream ss(argv[2]);
	int hfn_cap; // number of hash functions 
	if (!(ss >> hfn_cap)) {
		cerr << "Invalid number: " << argv[2] << '\n';
	} else if (!ss.eof()) {
		cerr << "Trailing characters after number: " << argv[2] << '\n';
	}
	
	// create indexing system
	//clock_t t;
	//t = clock();
	Indexing system = Indexing(hfn_cap);
	//t = clock() - t;
	//printf("The indexing system took %f seconds to build\n", ((double)t)/CLOCKS_PER_SEC);

	// handle csv file 
	ifstream ifs;
	ifs.open(filename);

	// error message if file does not exist in directory
    if (!ifs) {
        cout << "error opening file " << filename << endl;
        return 1;
    }
	
	string line, mh_val, record;
	
	istringstream ss_records_cap(argv[5]); // number of minhash vals taken in
	int records_cap;
	if (!(ss_records_cap >> records_cap)) {
		cerr << "Invalid number: " << argv[4] << '\n';
	} else if (!ss_records_cap.eof()) {
		cerr << "Trailing characters after number: " << argv[4] << '\n';
	}


	// read in csv file to indexing system.	
	//t = clock();
	int record_count = 0; // just in case num_records > records_cap in the file..
    while (!ifs.eof() && record_count < records_cap) {
		getline(ifs, line);
		stringstream s(line);
		int hf = 0; // hash function number

		//read each column
		getline(s, record, ',');
		while (getline(s, mh_val, ',') && hf < hfn_cap){ // cap the number of minhash values being take in 
			if (mh_val != ""){
				system.add_record(hf, record, mh_val);
			}
			hf++;
		}
		record_count += 1;
    }
	
	//t = clock()-t;
	//printf("It took %f seconds to read in the csv file\n", ((double)t)/CLOCKS_PER_SEC);


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
	//t = clock();
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

				system.next_record(hf, nrecord, nmh_val, similar); // add record to respective vector

			}
			hf++;
		}
    }
	//t = clock() - t;
	//printf("It took %f seconds to find similar records to new record\n", ((double)t)/CLOCKS_PER_SEC);

    // now that we have the vector of vectors of similar records, get the top K similar! 
    vector<string> results; // initiate the final array of most similar records


	// The code below is for some vector 
	clock_t t = clock();
	istringstream ss_k(argv[4]);
	int tk;
	if (!(ss_k >> tk)) {
		cerr << "Invalid number: " << argv[4] << '\n';
	} else if (!ss_k.eof()) {
		cerr << "Trailing characters after number: " << argv[4] << '\n';
	}

	TopKElements top_k;
	results = top_k.topKFrequent(similar, tk);


	// print out the top K results!
	for (int i = 0; i < results.size(); i++) {
		if (i == results.size() - 1)
			cout << results[i] << endl;
		else
			cout << results[i] << ",";
	}


	// only thing to print:
	t = clock() - t;
	printf("%f\n", ((double)t)/CLOCKS_PER_SEC);
	//printf("It took %f seconds to find the top %d similar\n", ((double)t)/CLOCKS_PER_SEC, tk);




	return 0;
}
