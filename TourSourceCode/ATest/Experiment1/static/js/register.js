
Email.onchange = function(){
    var Email = this.value;
    var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    if(!reg.test(Email)){
        alert("邮箱格式不正确，请重新输入！");
        return false;
    }
}

username.onchange = function(){
    var  username = this.value;
    var reg = /^\S{1,5}$/;
    if(!reg.test( username)){
        alert("用户名长度不能超过5位！");
        return false;
    }
}

password.onchange = function(){
    var password = this.value;
    var reg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,}$/;
    if(!reg.test(password)){
        alert("密码长度要大于6位，由数字和字母组成,请重新输入！");
        return false;
    }
}


function jump(){
    localStorage.setItem("username",document.getElementById("username").value);
    localStorage.setItem("password",document.getElementById("password").value);
    localStorage.setItem("email",document.getElementById("Email").value);

    var Email = document.getElementById("Email");
    var User = document.getElementById("username");
    var Password = document.getElementById("password");

    if(Email.value === "" || User.value ==="" || Password.value ==="") {
        alert("邮箱、用户名、或密码不能为空！");
        return false;
    }
    else {
        alert("注册成功，欢迎进入登录界面!!!");
        window.location.href="http://localhost:9999";
    }

}

function jump3(){
    window.location.href="login.html";
}
