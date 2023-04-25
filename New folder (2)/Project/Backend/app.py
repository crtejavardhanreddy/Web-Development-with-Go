from flask import Flask
from flask import request
app= Flask(__name__)
@app.route('/login',methods = ['POST','GET'])
def login():
    if request.method=='POST':
        user=request.form['userName']
        password=request.form['userPassWord']
        if user=='Teja' and password=='7483':
            return 'Welcome %s' %user
        else:
            return 'login failed %s' %user
    else:
        user=request.args.get('userName')
        password=request.args.get('userPassWord')
        if user=='Teja' and password=='7483':
            return 'Welcome %s' %user
        else:
            return 'login failed %s' %user
    
if __name__=="__main__":
    app.run(debug=True)
    