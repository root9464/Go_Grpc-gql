
source ./db.sh

all_run(){
  start_dbpsq &&
  go run ../core/main.go &&
  go run ../services/auth/auth.go
}


if [ $# -eq 1 ]; then
    "$1"
fi