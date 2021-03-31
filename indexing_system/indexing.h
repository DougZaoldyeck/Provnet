#ifndef INDEXING_H
#define INDEXING_H

#include <string>
#include <unordered_map>
#include <list>
#include <vector>
#include "sha256.h"
using namespace std;


class Indexing
{
protected:
	int k;
	vector<unordered_map<string, list>> hash_functions;
	// TODO why are we using a linked list and not a set? order matters?

public:
	Indexing();
	Indexing(int);
	~Indexing();
	// TODO what data types will the `r` be?
	void add_record(int n, string r, string mh_val);
};

#endif
