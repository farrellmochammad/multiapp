import mysql.connector as mysql
import os

def getDb():
    # enter your server IP address/domain name
    HOST = os.environ['HOST'] # or "domain.com"
    # database name, if you want just to connect to MySQL server, leave it empty
    DATABASE = os.environ['DATABASE']
    # this is the user you create
    USER = os.environ['USER']
    # user password
    PASSWORD = os.environ['PASSWORD']
    # connect to MySQL server
    db_connection = mysql.connect(host=HOST, database=DATABASE, user=USER, password=PASSWORD)
    print("Connected to:", db_connection.get_server_info())

    return db_connection
