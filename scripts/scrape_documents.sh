cd ..
mkdir -p books
for counter in $(seq 1 $1); do
    curl -s "https://www.gutenberg.org/cache/epub/$counter/pg$counter.txt" > books/pg$counter.txt
    mkdir -p books/pg
    split -a 4 -d -l 25 --additional-suffix=.txt books/pg$counter.txt books/pg/doc$counter
done
