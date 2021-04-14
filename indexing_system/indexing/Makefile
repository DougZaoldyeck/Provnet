CC=			g++
CFLAGS=		-std=c++11

index:		index.o indexing.o sha256.o
		$(CC) $(CFLAGS) $^ -o $@ 

sha256.o:	sha256.cpp
		$(CC) -c $(CFLAGS) $^ -o $@

index.o:	index.cpp
		$(CC) -c $(CFLAGS) $^ -o $@

indexing.o:	indexing.cpp 
		$(CC) -c $(CFLAGS) $^ -o $@

clean:		
		rm -f *.o index
