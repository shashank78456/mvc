function signupHandler(){
    document.getElementById("signup").addEventListener("submit", async (e)=>{

        e.preventDefault();
        const name = document.getElementById("name").value.trim();
        const username = document.getElementById("username").value.trim();
        const password = document.getElementById("password").value.trim();
        const cpassword = document.getElementById("cpassword").value.trim();
        
        if(name.length!=0 && username.length!=0 && password.length!=0 && cpassword.length!=0) {
            if(await checkStrength(password)) {
                if(cpassword===password) {
                    const response = await post({name: name, username: username, password: password},`http://localhost:3000/signup`);
                    const res = await response.json();
                    if(res.isValid) {
                        window.location.href = `http://localhost:3000/${res.userType}/home`;
                    }
                    else {
                        window.alert("User already exists");
                    }
                }
                else{
                    window.alert("Password entered is different");
                }
            }
            else{
                window.alert("Weak password. Must have length more than 7 and contain atleast 1 number and 1 special character");
            }
        }
        else {
            window.alert("Please Enter Valid Information");
        }

    })
}

async function post(data, url) {
    return new Promise((resolve) => {
        fetch(url, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
        .then(response => {
            resolve(response);
        })
        .catch(error => {
            console.error(error);
        });
    });
}

async function get(url) {
    return new Promise((resolve) => {
        fetch(url)
        .then(response => {
            resolve(response);
        })
        .catch(error => console.error(error));
    });
}

async function checkStrength(password) {
    let isStrong = false;
    return new Promise((resolve) => {
        if(password.length>=8) {
            let specialCount=0;
            let numberCount=0;
            for(let i=0; i<password.length; i++) {
                let ord = password.charCodeAt(i);
                if(ord>=48 && ord<=57)
                    numberCount++;
                else if(!(ord>=65 && ord<=90) && !(ord>=97 && ord<=122))
                    specialCount++;
                if(numberCount>=1 && specialCount>=1){
                    isStrong = true;
                    break;
                }
            }
        }
        resolve(isStrong);
    });
}

signupHandler();