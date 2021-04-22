CC=			g++
CFLAGS=		-std=c++11

main:		main.o topK.o indexing.o sha256.o
		$(CC) $(CFLAGS) $^ -o $@

sha256.o:	sha256.cpp
		$(CC) -c $(CFLAGS) $^ -o $@

main.o:		main.cpp
		$(CC) -c $(CFLAGS) $^ -o $@

topK.o:		topK.cpp
		$(CC) -c $(CFLAGS) $^ -o $@

indexing.o:	indexing.cpp 
		$(CC) -c $(CFLAGS) $^ -o $@

clean:		
		rm -f *.o main
