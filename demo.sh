#!/usr/bin/env sh

#set -x # Uncomment if you want to see all commands run

SERVER=bin/server
CLIENT=bin/client

RED='\033[0;31m'
NC='\033[0m' # No Color

ACTION=1
TOTAL_ACTIONS=13

red() {
    printf "${RED}($ACTION/$TOTAL_ACTIONS) $1${NC}\n"
    ((ACTION++))
}

#set -e
make all
make db

if [ ! -e server.crt ]; then
    red "sorry ya gotta do this once -- just hit enter for everything except the \"Common Name\"; for that one use \"localhost\""

    openssl req -new -key server.key -out server.csr
    openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
    ((TOTAL_ACTIONS++))
fi

red "Start server"
pkill -f bin/server
nohup $SERVER &

# delay to let server start
sleep 1

red "Try to get secret"
$CLIENT secret asdf

red "Create a user -- please type in a password"
$CLIENT useradd testuser

red "Now get a token"
$CLIENT auth testuser

red "Please copy and paste the new token above:"
read TOKEN

red "Now let's try to get that secret"
$CLIENT secret $TOKEN

red "That was a nice secret. Now, your password is bad. Let's change it."
$CLIENT passwd testuser $TOKEN

red "Logout the first session:"
$CLIENT logout $TOKEN

red "Prove we're logged out:"
$CLIENT secret $TOKEN

red "Start a new session with the new password"
$CLIENT auth testuser

red "Please copy and past new token again:"
read TOKEN

red "Confirm we can login with the new session"
$CLIENT secret $TOKEN

red "That's a wrap, folks."

pkill -f bin/server

