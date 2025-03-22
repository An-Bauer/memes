function addRegister() {
    let register = document.getElementById("register")
    let username = document.getElementById("regUsername")
    let password1 = document.getElementById("regPassword1")
    let password2 = document.getElementById("regPassword2")
    let msg = document.getElementById("regMsg")
    let reg = document.getElementById("regBut")

    register.addEventListener("input", function (event) {
        let error = false
        let errorMsg = ""

        password2.classList.remove('wrong');
        if (password1.value != password2.value) {
            error = true
            if (password2.value != "") {
                password2.classList.add('wrong');
                errorMsg = "Passwörter ungleich"
            }
        }

        password1.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,20}$/).test(password1.value)) {
            error = true
            if (password1.value != "") {
                password1.classList.add('wrong');
                errorMsg = "Passwort ungültig"
            }
        }

        username.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,10}$/).test(username.value)) {
            error = true
            if (username.value != "") {
                username.classList.add('wrong');
                errorMsg = "Benutzername ungültig"
            }
        }

        reg.disabled = error
        console.log(error)
        msg.innerHTML = errorMsg
    });
} addRegister()


function addLogin() {
    let login = document.getElementById("login")
    let username = document.getElementById("logUsername")
    let password = document.getElementById("logPassword")
    let msg = document.getElementById("logMsg")
    let reg = document.getElementById("logBut")

    login.addEventListener("input", function (event) {
        let error = false
        let errorMsg = ""

        password.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,20}$/).test(password.value)) {
            error = true
            if (password.value != "") {
                password.classList.add('wrong');
                errorMsg = "Passwort ungültig"
            }
        }

        username.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,10}$/).test(username.value)) {
            error = true
            if (username.value != "") {
                username.classList.add('wrong');
                errorMsg = "Benutzername ungültig"
            }
        }

        reg.disabled = error
        msg.innerHTML = errorMsg
    });
} addLogin()


function regFunc(e) {
    const formData = new URLSearchParams();
    formData.append("user", document.getElementById("regUsername").value);
    formData.append("password", document.getElementById("regPassword1").value);

    fetch("/api/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: formData.toString()
    })
        .then(response => console.log(response.text().then()))
        .catch(error => console.error("Error:", error));

} document.getElementById("regBut").addEventListener("click", regFunc)

function logFunc(e) {
    const formData = new URLSearchParams();
    formData.append("user", document.getElementById("logUsername").value);
    formData.append("password", document.getElementById("logPassword").value);

    fetch("/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: formData.toString()
    })
        .then(data => console.log(data))
        .catch(error => console.error("Error:", error));

} document.getElementById("logBut").addEventListener("click", logFunc)

function closeAccount(e) {
    if (e.target == this) {
        document.getElementById("bgAccount").hidden = true
    }
} document.getElementById("bgAccount").addEventListener("click", closeAccount)

function showAccount(e) {
    document.getElementById("bgAccount").hidden = false
} document.getElementById("account").addEventListener("click", showAccount)