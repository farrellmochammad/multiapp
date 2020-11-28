import mysql.connector as mysql

def getDb():
    # enter your server IP address/domain name
    HOST = "indikost.com" # or "domain.com"
    # database name, if you want just to connect to MySQL server, leave it empty
    DATABASE = "u1107404_efishery"
    # this is the user you create
    USER = "u1107404_efishery"
    # user password
    PASSWORD = "efishery123!"
    # connect to MySQL server
    db_connection = mysql.connect(host=HOST, database=DATABASE, user=USER, password=PASSWORD)
    print("Connected to:", db_connection.get_server_info())

    return db_connection
