function adminHandler(){

    document.getElementById("catalog").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/admin/home`;
    })

    document.getElementById("requests").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/admin/requests`;
    })

    document.getElementById("logout").addEventListener("click", async ()=> {
        document.cookie = 'token=; Path=/; Expires=Thu, 01 Jan 2000 00:00:01 GMT;';
        window.location.href = `http://localhost:3000`;
    })

    try {
        document.getElementById("prompt-open").addEventListener("click", async ()=> {
            window.location.href = `http://localhost:3000/admin/add`;
        })
    }
    catch{};

    try {
    document.getElementById("adreq").addEventListener("click", async ()=> {
        window.location.href = `http://localhost:3000/admin/adreq`;
    })
    }
    catch{};

    try {
    document.getElementById("prompt-form").addEventListener("submit", async (e)=> {
        e.preventDefault();
        const bookName = document.getElementById("new-book").value.trim();
        const author = document.getElementById("new-author").value.trim();
        if(bookName.length!=0) {
            const response  = await post({book: bookName, author: author}, `http://localhost:3000/admin/add_new_book`);
            const res = await response.json();
            if(res.isAdded) {
                window.alert("Added Succesfully");
                window.location.href = `http://localhost:3000/admin/home`;
            }
            else {
                window.alert("Book already exists");
            }
        }
        else {
            window.alert("Enter Valid Book Name");
        }
    })
    }
    catch{};

    try {
    const addButton = document.getElementsByClassName("add");
    for(let i=0; i<addButton.length; i++) {
        addButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            await post({book: addButton[i].value, toDo: "add"}, `http://localhost:3000/admin/add_book`);
            window.alert("Added Successfully");
            window.location.href = `http://localhost:3000/admin/home`;
        })
    }
    
    const removeButton = document.getElementsByClassName("remove");
    for(let i=0; i<removeButton.length; i++) {
        removeButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            await post({book: removeButton[i].value, toDo: "remove"}, `http://localhost:3000/admin/remove_book`);
            window.alert("Removed Successfully");
            window.location.href = `http://localhost:3000/admin/home`;
        })
    }

    const deleteButton = document.getElementsByClassName("delete");
    for(let i=0; i<deleteButton.length; i++) {
        deleteButton[i].addEventListener("click", async (e)=> {
            e.preventDefault();
            const response = await post({book: deleteButton[i].value, toDo: "delete"}, `http://localhost:3000/admin/delete_book`);
            const res = await response.json();
            if(res.isDeleted) {
                window.alert("Deleted Successfully");
                window.location.href = `http://localhost:3000/admin/home`;
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
        let requestid = acceptButton[i].value;
        acceptButton[i].addEventListener("click", async ()=> {
            await post({requestid: requestid}, `http://localhost:3000/admin/accept_request`);
            window.alert("Accepted Successfully");
            window.location.href = `http://localhost:3000/admin/requests`;
        })
    }
    }
    catch{};

    try {
    const acceptButtonR = document.getElementsByClassName("acceptR");
    const denyButtonR = document.getElementsByClassName("denyR");
    for(let i=0; i<acceptButtonR.length; i++) {
        acceptButtonR[i].addEventListener("click", async ()=> {
            await post({user: acceptButtonR[i].value}, `http://localhost:3000/admin/accept_admin`);
            window.alert("Accepted Successfully");
            window.location.href = `http://localhost:3000/admin/adreq`;
        })
    }
    
    for(let i=0; i<denyButtonR.length; i++) {
        denyButtonR[i].addEventListener("click", async ()=> {
            await post({user: denyButtonR[i].value}, `http://localhost:3000/admin/deny_admin`);
            window.alert("Denied Successfully");
            window.location.href = `http://localhost:3000/admin/adreq`;
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

adminHandler();
