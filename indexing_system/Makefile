CC=			g++
CFLAGS=		-std=c++11

index:		index.o indexing.o
		$(CC) $(CFLAGS) $^ -o $@ 

index.o:	index.cpp indexing.h
		$(CC) -c $(CFLAGS) index.cpp -o $@

indexing.o:	indexing.cpp indexing.h
		$(CC) -c $(CFLAGS) indexing.cpp -o $@

clean:		
		rm -f *.o index
