from flask import Flask, request, jsonify, Blueprint
from auth import authusecase
import middleware

app = Flask(__name__)
v1 = Blueprint("api_V1", __name__)



@v1.route("/register", methods=['POST'])
def register():
    req_data = request.get_json(force=True)
  
    user = {
        "phone" : req_data['phone'],
        "name" : req_data['name'],
        "role" : req_data['role']
    }

    authUsecase = authusecase.auth_usecase()
    password = authUsecase.insertUser(user)

    return jsonify (
        status = "success",
        password = password
    )

@v1.route("/login", methods=['POST'])
def login():
    req_data = request.get_json(force=True)
    user = {
        "phone" : req_data['phone'],
        "password" : req_data['password']
    }

    authUsecase = authusecase.auth_usecase()
    status,token = authUsecase.getUserByPhone(user)
    if (status) :
        return jsonify (
            status = "success",
            token = token
        )
    else :
        return jsonify (
            status = "false",
            message = "Username atau password tidak valid"
        )

@v1.route("/userinfo", methods=['POST'])
def userinfo():
    header = request.headers.get('Authorization')
    authUsecase = authusecase.auth_usecase()
    status,message = authUsecase.getUserInfoJwt(header)
    if status :
        return jsonify (
            phone = message["phone"],
            name = message["name"],
            role = message["role"],
            timestamp = message["timestamp"],
        )
    else :
        return jsonify (
            stats = "failed",
            message = message
        )

@v1.route("/area", methods=['POST'])
def area():
    print("Hello")

        
    
    

if __name__ == "__main__":
    port = 8015
    app.register_blueprint(v1, url_prefix="/api/v1")
    app.run(host='0.0.0.0', port=port)