async function register() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const email = document.getElementById('email').value;

    const response = await fetch("/register", {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"username": username, "password": password, "email": email}),
    });

    const data = await response.json();
    console.log(data);
    //localStorage.setItem('token', data.token);
    if(response.ok) {
        window.location.href = "/login";
    }
}

async function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const response = await fetch("/login", {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"username": username, "password": password}),
    });

    const data = await response.json();
    console.log(data);
    if(response.ok) {
        window.location.href = "/index";
    }
    //localStorage.setItem('token', data.token);

}

async function getProfile() {
    const role = {0: "admin", 1: "student", 2: "Guest"}

    const response = await fetch("/user", {
        method: 'POST',
    });
    const data = await response.json()
    console.log(data);

    if(!response.ok){
        window.location.href = data["redirect"];
        // const prompt = document.getElementById("prompt");
        // prompt.innerText = data["error"];
        // prompt.style.color = "red";
    } else {
        document.getElementById("username").value = data["username"];
        document.getElementById("email").value = data["email"];
        document.getElementById("role").value = role[data["role"]];
    }
}

async function updateProfile() {
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const response = await fetch("/user", {
        method: 'PUT',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"username": username, "email": email}),
    })

    const data = await response.json();
    console.log(data);
    if(response.ok){
        window.location.href = "/user";
    }
}

async function changePassword() {
    const originalPwd = document.getElementById('originalPwd').value;
    const newPwd = document.getElementById('newPwd').value;
    const response = await fetch("/user", {
        method: 'PATCH',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"originalPwd": originalPwd, "newPwd": newPwd}),
    })

    const data = await response.json();
    console.log(data);
    if(response.ok){
        window.location.href = "/login";
    }
}

async function logout() {
    const response = await fetch("/logout", {
        method: 'POST',
    })

    const data = await response.json();
    console.log(data);
    if(response.ok){
        window.location.href = "/index";
    }
}

async function deleteUser() {
    const response = await fetch("/user", {
        method: 'DELETE',
    })

    const data = await response.json();
    console.log(data);
    if(response.ok){
        window.location.href = "/register";
    }
}