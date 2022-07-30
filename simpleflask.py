import sys
import time


# simulate flask route
class SimpleFlask:
    def __init__(self):
        self.router = {}
    def register(self, path):
        def decorator(h):
            self.router[path] = h
            print("I have register func \"%s\" to path \"%s\""% (h.__name__, path))
            # just return the original handler
            return h
        return decorator

app = SimpleFlask()


# tag:section1 
# def hello():
#     print("hello!")
# hello = app.register("/home")(hello)

# syntax sugar for section1
@app.register("/home")
def hello(params):
    print("hello!", params)

def trace(h):
    def inner():
        print("execute h")
        h()
        print("end of h")
    return inner

@trace
def handler():
    print("this is the handler")

if __name__ == '__main__':
    """python3 -i test.py "/home" "/index/" "random"

        I have register func "hello" to path "/home"
        ['test.py', '/home', '/index/', 'random']
        arg =  test.py
        arg =  /home
        hello! 1659186959.8527172
        arg =  /index/
        arg =  random
    """
    print(sys.argv)
    for arg in sys.argv:
        print("arg = ", arg)
        if arg  in app.router:
            app.router[arg](time.time())

