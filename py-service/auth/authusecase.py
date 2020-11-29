from auth import authrepo
from passlib.hash import bcrypt
import random
import string
import jwt 
import datetime
import os

class auth_usecase:
    def insertUser(self,user):
        authRepo = authrepo.auth_repository()
        password = self.getRandomChar(4)
        hashed = bcrypt.hash(password)
        user["password"] = hashed
        if authRepo.insertUser(user) == False :
            return False
        else:
            return password

    def getUserByPhone(self,user):
        authRepo = authrepo.auth_repository()
        userVerify = authRepo.getUserByPhone(user)
        if bcrypt.verify(user["password"], userVerify["password"]) :
            encoded_jwt = jwt.encode({
                'phone' : userVerify["phone"],
                'name' : userVerify["name"],
                'role' : userVerify["role"],
                'password' : userVerify["password"],
                'timestamp' : datetime.datetime.now().timestamp(),
            },os.environ['SIGNING_KEY'],algorithm='HS256')
            return True,encoded_jwt.decode("utf-8")
        else :
            return False,None

    def getUserInfoJwt(self,header):
        if ("Bearer: " not in header) :
            return False, "login failed, token not valid"
        else :
            headerString = self.convertToString(header) 
            arrHeader = self.convertToList(headerString)
            if (len(arrHeader) > 2) :
                return False, "Authorization not valid"
            token = arrHeader[1]
            try:
                user = jwt.decode(token,os.environ['SIGNING_KEY'], algorithms=['HS256'], verify= True)
                return True,user
            except :
                return False,"Token expired, please login back"

        
        
    def getRandomChar(self,length):
        letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345679"
        result_str = ''.join(random.choice(letters) for i in range(length))
        return result_str

    def splitToWord(self,listword):
        return ' '.join(listword).split() 

    def convertToString(self,listword):  
        word = "" 
        return (word.join(listword)) 

    def convertToList(self,string): 
        li = list(string.split(" ")) 
        return li 

