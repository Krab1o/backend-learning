'use strict';

let user = {
  email: undefined,
  password: undefined,
}

let button = document.querySelector(".btn");

button.addEventListener("click", (e) => {
  e.preventDefault();
  const data = new FormData(document.getElementById("loginForm"));
  loginAttempt(data);
})

async function loginAttempt(formData) {
  user.email = formData.get("emailInput");
  user.password = formData.get("passwordInput");
  try {
    const loginUrl = "http://localhost:8080/api/login";
    let response = await fetch(loginUrl, {
      method: "POST",
      body: JSON.stringify(user),
    });
    if (response.status == 401) {
      showWrongPassword();
      return;
    }
    else {
      proceedRequest();
      return;
    }
  } catch (err) {
    alert(err)
    return;
  }
}

function showWrongPassword() {
  let wrongPassword = document.getElementById("wrongPassword");
  if (getComputedStyle(wrongPassword).display === "none") {
    wrongPassword.style.display = "block";
    let passwordForm = document.getElementById("passwordForm");
    passwordForm.style.marginBottom = 0;
  }
}

async function proceedRequest() {
  location.href = "http://localhost:8080/articles"
}