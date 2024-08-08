function clientHandler() {
  document.getElementById("reqadmin").addEventListener("click", async () => {
    const response = await post(
      { hasAdminRequest: true },
      `http://localhost:3000/client/admin_request`
    );
    const res = await response.json();
    if (!res.hasAlreadyRequested) {
      window.alert("Successfully Requested to Become Admin");
    } else {
      window.alert("Already Requested to Become Admin");
    }
  });

  document.getElementById("view").addEventListener("click", async () => {
    window.location.href = `http://localhost:3000/client/home`;
  });

  document.getElementById("return").addEventListener("click", async () => {
    window.location.href = `http://localhost:3000/client/return`;
  });

  document.getElementById("borrow").addEventListener("click", async () => {
    window.location.href = `http://localhost:3000/client/history`;
  });
  document.getElementById("logout").addEventListener("click", async () => {
    document.cookie =
      "accesstoken=; Path=/; Expires=Thu, 01 Jan 2000 00:00:01 GMT;";
    window.location.href = `http://localhost:3000`;
  });

  try {
    const borrowBooks = document.getElementsByClassName("borrow");
    for (let i = 0; i < borrowBooks.length; i++) {
      borrowBooks[i].addEventListener("click", async () => {
        const response = await post(
          { bookid: parseInt(borrowBooks[i].value) },
          `http://localhost:3000/client/request_book`
        );
        const res = await response.json();
        if (res.hasAlreadyRequested) {
          window.alert("Requested Successfully");
          window.location.href = `http://localhost:3000/client/home`;
        } else {
          window.alert("You have already requested for the Book");
        }
      });
    }
  } catch {}

  try {
    const returnBooks = document.getElementsByClassName("return");
    for (let i = 0; i < returnBooks.length; i++) {
      returnBooks[i].addEventListener("click", async () => {
        await post(
          { bookid: parseInt(returnBooks[i].value) },
          `http://localhost:3000/client/return_book`
        );
        window.alert("Returned Successfully");
        window.location.href = `http://localhost:3000/client/return`;
      });
    }
  } catch {}
}

async function post(data, url) {
  return new Promise((resolve) => {
    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        console.error(error);
      });
  });
}

async function get(url) {
  return new Promise((resolve) => {
    fetch(url)
      .then((response) => {
        resolve(response);
      })
      .catch((error) => console.error(error));
  });
}

clientHandler();
