function superadminHandler(){

    document.getElementById("catalog").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/superadmin/home`;
    })

    document.getElementById("requests").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/superadmin/requests`;
    })

    document.getElementById("logout").addEventListener("click", async ()=> {
        document.cookie = 'accesstoken=; Path=/; Expires=Thu, 01 Jan 2000 00:00:01 GMT;';
        window.location.href = `http://localhost:3000`;
    })

    try {
        document.getElementById("prompt-open").addEventListener("click", async ()=> {
            window.location.href = `http://localhost:3000/superadmin/add`;
        })
    }
    catch{};

    try {
    document.getElementById("adreq").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/superadmin/adreq`;
    })
    }
    catch{};

    try {
    document.getElementById("prompt-form").addEventListener("submit", async (e)=> {
        e.preventDefault();
        const bookname = document.getElementById("new-book").value.trim();
        const author = document.getElementById("new-author").value.trim();
        const quantity = document.getElementById("new-quantity").value.trim();
        if(quantity>=0) {
            if(author.length!=0) {
                if(bookname.length!=0) {
                    const response  = await post({bookname: bookname, author: author, quantity: parseInt(quantity)}, `http://localhost:3000/superadmin/add_new_book`);
                    const res = await response.json();
                    if(res.isAdded) {
                        window.alert("Added Succesfully");
                        window.location.href = `http://localhost:3000/superadmin/home`;
                    }
                    else {
                        window.alert("Book already exists");
                    }
                }
                else {
                    window.alert("Enter Valid Book Name");
                }
            }
            else {
                window.alert("Enter Valid Author Name");
            }
        }
        else {
            window.alert("Quantity should not be negative");
        }
    })
    }
    catch{};

    try {
    const addButton = document.getElementsByClassName("add");
    for(let i=0; i<addButton.length; i++) {
        addButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            await post({bookid: parseInt(addButton[i].value), toDo: "add"}, `http://localhost:3000/superadmin/add_book`);
            window.location.href = `http://localhost:3000/superadmin/home`;
        })
    }
    
    const removeButton = document.getElementsByClassName("remove");
    for(let i=0; i<removeButton.length; i++) {
        removeButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            await post({bookid: parseInt(removeButton[i].value), toDo: "remove"}, `http://localhost:3000/superadmin/remove_book`);
            window.location.href = `http://localhost:3000/superadmin/home`;
        })
    }

    const deleteButton = document.getElementsByClassName("delete");
    for(let i=0; i<deleteButton.length; i++) {
        deleteButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            const response = await del({bookid: parseInt(deleteButton[i].value), toDo: "delete"}, `http://localhost:3000/superadmin/delete_book`);
            const res = await response.json();
            if(res.isDeleted) {
                window.alert("Deleted Successfully");
                window.location.href = `http://localhost:3000/superadmin/home`;
            }
            else{
                window.alert("Cannot delete this book as it is already borrowed");
            }
        })
    }

    }
    catch{};

    try {
    const acceptButton = document.getElementsByClassName("accept");
    for(let i=0; i<acceptButton.length; i++) {
        let requestid = parseInt(acceptButton[i].value);
        acceptButton[i].addEventListener("click", async ()=> {
            await post({requestid: requestid}, `http://localhost:3000/superadmin/accept_request`);
            window.alert("Accepted Successfully");
            window.location.href = `http://localhost:3000/superadmin/requests`;
        })
    }
    }
    catch{};

    try {
    const acceptButtonR = document.getElementsByClassName("acceptR");
    const denyButtonR = document.getElementsByClassName("denyR");
    for(let i=0; i<acceptButtonR.length; i++) {
        acceptButtonR[i].addEventListener("click", async ()=> {
            await post({userid: parseInt(acceptButtonR[i].value)}, `http://localhost:3000/superadmin/accept_admin`);
            window.alert("Accepted Successfully");
            window.location.href = `http://localhost:3000/superadmin/adreq`;
        })
    }
    
    for(let i=0; i<denyButtonR.length; i++) {
        denyButtonR[i].addEventListener("click", async ()=> {
            await post({userid: parseInt(denyButtonR[i].value)}, `http://localhost:3000/superadmin/deny_admin`);
            window.alert("Denied Successfully");
            window.location.href = `http://localhost:3000/superadmin/adreq`;
        })
    }
    }
    catch{};
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

async function del(data, url) {
    return new Promise((resolve) => {
        fetch(url, {
            method: "DELETE",
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

superadminHandler();
