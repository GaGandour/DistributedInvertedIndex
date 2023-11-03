cd ..
mkdir -p books
for counter in $(seq 1 $1); do curl https://www.gutenberg.org/cache/epub/$counter/pg$counter.txt > books/pg$counter.txt; done