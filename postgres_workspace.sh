sudo apt update
yes | sudo apt install postgresql postgresql-contrib
echo "init script"
sudo service postgresql start
echo "inited psql"
sudo -u postgres createuser person_master -s -d
sudo -u postgres createdb person_register --owner=person_master
sudo -u postgres psql -d person_register -f init_ddbb_users.sql

echo "end db"
go get -u github.com/lib/pq
