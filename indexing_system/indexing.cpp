/* Class for indexing system */

#include <iostream>
#include "indexing.h"

Indexing::Indexing(int k) {
	hash_functions.reserve(k);
}

Indexing::Indexing() {
	hash_functions.reserve(1200);
}

Indexing::~Indexing() { }

void Indexing::add_record(int n, string r, string mh_val) {
	//TODO do a search, then add it to the back.
	// others = search(n, mh_val);
	hash_functions.at(n).at(mh_val).push_back(r);
	// return others;
}
// TODO is Record supposed to be a string representing the actual record?
// i.e. what data type?
// TODO why are we using a linked list and not a set? does the order really matteR? 
// TODO 
