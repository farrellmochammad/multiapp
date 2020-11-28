import db
import json

class auth_repository:
    def getUserByPhone(self,user):
        conn = db.getDb()
        sql_select_Query = "select phone,name,role,password from Users where phone = " + "'" + user["phone"] + "'"
        cursor = conn.cursor()
        cursor.execute(sql_select_Query)
        records = cursor.fetchall()

        for row in records:
            user = {}
            user["phone"] = row[0]
            user["name"] = row[1]
            user["role"] = row[2]
            user["password"] = row[3]
            return user
        
    
    def insertUser(self,user):
        sql = "INSERT INTO Users (phone,name,role,password) VALUES (%s,%s,%s,%s)"
        val = (user["phone"],user["name"],user["role"],user["password"])
        conn = db.getDb()
        cursor = conn.cursor()
        cursor.execute(sql,val)

        conn.commit()
        
        return user
    