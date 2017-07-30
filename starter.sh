# grab directory
SCRIPT=`python -c "import os,sys; print(os.path.realpath(os.path.expanduser(sys.argv[1])))" "${BASH_SOURCE:-$0}"`
DIR=$(dirname $(dirname $SCRIPT))

# make temp compose directory
sudo mkdir compose
cd $DIR/compose

# clone relevant git repo
git clone https://github.com/openzipkin/docker-zipkin.git
cd docker-zipkin 
mv docker-compose.yml $DIR/docker-compose.yml 

# run script to add in our testServer to the compose file
./update_compose.sh

# cleanup the repo
cd ..
rm -rf compose
