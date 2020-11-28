def Middleware(func):
    def handler(request, *args, **kwargs):
        try:
            print("Hello before")
            return func(request, *args, **kwargs)
        except :
            return "son, i am dissapoint"
    return handler