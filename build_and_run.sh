RED=`tput setaf 1`
GREEN=`tput setaf 2`
NC=`tput sgr0`
BROWN=`tput setaf 5`

exec_command() {
	echo "Executing: $2	"
	sleep 2
	$2
	if [ $? -ne 0 ]; then
		echo "${RD} $1 Failed"; exit 1
	fi
	echo "${BROWN} $1 ${GREEN} SUCCESS ${NC}"
}

exec_command "Getting Code from Git" "git clone https://github.com/skoyya/grpc_demo.git"
exec_command "Changing to workspace" "cd grpc_demo"
export GOPATH=`pwd`
exec_command "Getting GRPC" "go get -u google.golang.org/grpc"
exec_command "Building server " "go build -i -v -o bin/server"
exec_command "Building client"  "go build -i -v -o bin/client client"
exec_command "Running unit-test on server" "go test -v server"
sleep 1
bin/server &
echo "--------- Running the server --------------${GREEN} SUCCESS ${NC}"
sleep 5
echo "--------- Running the client --------------"
sleep 2
bin/client 
sleep 2
