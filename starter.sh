# grab directory
SCRIPT=`python -c "import os,sys; print(os.path.realpath(os.path.expanduser(sys.argv[1])))" "${BASH_SOURCE:-$0}"`
DIR=$(dirname $SCRIPT)

# clone relevant git repo
git clone https://github.com/openzipkin/docker-zipkin.git
cd docker-zipkin
mv docker-compose.yml $DIR/docker-compose.yml

# cleanup the repo
cd ..
rm -rf $DIR/docker-zipkin
