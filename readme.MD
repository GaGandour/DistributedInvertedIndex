# DISTRIBUTED INVERTED INDEX

## What it is

TODO

## How to use it

### First step: Preparing the environment

First, we have to clone the repo. Make sure you have docker installed.

```
git clone https://github.com/GaGandour/DistributedInvertedIndex.git
```

From now on, all steps will be executed on the scripts folder:

```
cd scripts
```

### Second step: Scrape data.

We have to gather data to put in our inverted index:

```
sh scrape_documents.sh 20
```

The number 20 is related to the ammount of documents we will download. It can be any positive integer. The greater the number, more data will be downloaded.

This script downloads files from [https://www.gutenberg.org](https://www.gutenberg.org) and splits them into smaller files.

### Third step: build docker image.

After scraping the data, we can build the docker image:

```
sh build_image.sh
```

### Final step: execution.

```
sh start_docker 4
```

The number 4 is the number of workers that will be run. It can be any positive integer. 

**WARNING**: be careful with the number of workers you use. If your machine is not powerful, to many workers could cause a problem. Using 3 workers is a safe approach.

## Contributors:

+ Gabriel Gandour [@GaGandour](https://github.com/GaGandour)
+ Emmanuel da Silva Dias [@emmanuelsdias](https://github.com/emmanuelsdias)
+ Marcelo Buga [@BugaM](https://github.com/BugaM)
+ Gabriel Ribeiro Pardini [@gabrielribpar](https://github.com/gabrielribpar)
