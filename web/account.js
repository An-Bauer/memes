function addRegister() {
    let register = document.getElementById("register")
    let username = document.getElementById("regUsername")
    let password1 = document.getElementById("regPassword1")
    let password2 = document.getElementById("regPassword2")
    let msg = document.getElementById("regMsg")
    let reg = document.getElementById("regBut")

    register.addEventListener("input", function (event) {
        let error = false
        let msg = "registrieren"

        password2.classList.remove('wrong');
        if (password1.value != password2.value) {
            error = true
            if (password2.value != "") {
                password2.classList.add('wrong');
                msg = "Passwörter ungleich"
            }
        }

        password1.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,20}$/).test(password1.value)) {
            error = true
            if (password1.value != "") {
                password1.classList.add('wrong');
                msg = "Passwort ungültig"
            }
        }

        username.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,10}$/).test(username.value)) {
            error = true
            if (username.value != "") {
                username.classList.add('wrong');
                msg = "Benutzername ungültig"
            }
        }

        reg.disabled = error
        reg.innerHTML = msg
    });
} addRegister()

function addLogin() {
    let login = document.getElementById("login")
    let username = document.getElementById("logUsername")
    let password = document.getElementById("logPassword")
    let log = document.getElementById("logBut")

    login.addEventListener("input", function (event) {
        let error = false
        let msg = "anmelden"

        password.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,20}$/).test(password.value)) {
            error = true
            if (password.value != "") {
                password.classList.add('wrong');
                msg = "Passwort ungültig"
            }
        }

        username.classList.remove('wrong');
        if (!(/^[a-zA-Z0-9]{4,10}$/).test(username.value)) {
            error = true
            if (username.value != "") {
                username.classList.add('wrong');
                msg = "Benutzername ungültig"
            }
        }

        log.disabled = error
        log.innerHTML = msg
    });
} addLogin()

async function register(e) {
    try {
        const regBut = document.getElementById("regBut")

        const formData = new URLSearchParams();
        formData.append("user", document.getElementById("regUsername").value);
        formData.append("password", document.getElementById("regPassword1").value);
        const request = {
            method: "POST",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            body: formData.toString(),
        }
        const response = await fetch("/api/register", request)

        if (!response.ok) {
            regBut.innerHTML = "interner Fehler"
            console.log("server error")
            return
        }

        const text = await response.text()

        if (text == "user already exists") {
            regBut.innerHTML = "Benutzer existiert schon"
        } else if (text == "success") {
            window.location.reload();
        } else {
            regBut.innerHTML = "interner Fehler"
            console.log(text)
            console.log("invalid response")
        }
    } catch (error) {
        logBut.innerHTML = "kein Internet"
        console.log(error)
    }
} document.getElementById("regBut").addEventListener("click", register)

async function login(e) {
    try {
        const logBut = document.getElementById("logBut")

        const formData = new URLSearchParams();
        formData.append("user", document.getElementById("logUsername").value);
        formData.append("password", document.getElementById("logPassword").value);
        const request = {
            method: "POST",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            body: formData.toString(),
        }
        const response = await fetch("/api/login", request)

        if (!response.ok) {
            logBut.innerHTML = "interner Fehler"
            console.log("server error")
            return
        }

        const text = await response.text()

        if (text == "user dosn't exist") {
            logBut.innerHTML = "Benutzer existiert nicht"
        } else if (text == "wrong password") {
            logBut.innerHTML = "Passwort falsch"
        } else if (text == "success") {
            window.location.reload();
        } else {
            logBut.innerHTML = "interner Fehler"
            console.log("invalid response")
        }
    } catch (error) {
        logBut.innerHTML = "kein Internet"
        console.log(error)
    }
} document.getElementById("logBut").addEventListener("click", login)

function closeAccount(e) {
    if (e.target == this) {
        document.getElementById("bgAccount").hidden = true
    }
} document.getElementById("bgAccount").addEventListener("click", closeAccount)

function showAccount(e) {
    document.getElementById("bgAccount").hidden = false
} document.getElementById("account").addEventListener("click", showAccount)