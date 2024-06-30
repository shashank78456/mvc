function loginHandler(){

    document.getElementById("login").addEventListener("submit", async (e)=>{

        e.preventDefault();
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        if(username.length!=0 && password.length!=0) {
            const response = await post({username: username, password: password},`http://localhost:3000`);
            const res = await response.json();
            if(res.isValid)
                window.location.href = `http://localhost:3000/${res.userType}/home`;
            else
                window.alert("Wrong credentials");
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

loginHandler();