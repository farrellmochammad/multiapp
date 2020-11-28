from functools import wraps
from flask import request

def validate_jwt(f):
    @wraps(f)
    def validate_function(*args, **kwargs):
        print("Header : ", request.headers.get("Authorization"))
        result = {}
        result["Hello"] = 123
        return result
      
    return validate_function

def splitToWord(listword):
    return ' '.join(listword).split() 

def convertToString(listword):  
    word = "" 
    return (word.join(listword)) 

def convertToList(string): 
    li = list(string.split(" ")) 
    return li 