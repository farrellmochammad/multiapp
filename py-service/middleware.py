from functools import wraps
from flask import request, jsonify, g
import jwt

def validate_jwt(f):
    @wraps(f)
    def validate_function(*args, **kwargs):
        header =  request.headers.get("Authorization")
        if ("Bearer: " not in header) :
            return jsonify (
                status = "failed",
                message = "login failed, token not valid"
            )
        else :
            headerString = convertToString(header) 
            arrHeader = convertToList(headerString)
            if (len(arrHeader) > 2) :
                return jsonify (
                    status = "failed",
                    message = "Authorization not valid"
                )
            token = arrHeader[1]
            try:
                user = jwt.decode(token,"efishery123", algorithms=['HS256'], verify= True)
                g.user = user
                return f(*args, **kwargs)
            except :
                return jsonify (
                    status = "failed",
                    message = "Token expired, please login back"
                )
      
    return validate_function

    

def splitToWord(listword):
    return ' '.join(listword).split() 

def convertToString(listword):  
    word = "" 
    return (word.join(listword)) 

def convertToList(string): 
    li = list(string.split(" ")) 
    return li 