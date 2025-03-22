
async function login() {
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
        } else if (text == "success") {
            reload()
        } else {
            logBut.innerHTML = "interner Fehler"
            console.log("invalid response")
        }
        throw new Error("invalid response");
    } catch (error) {
        logBut.innerHTML = "kein Internet"
        console.log(error)
    }
}