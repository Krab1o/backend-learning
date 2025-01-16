'use strict';

function loginAttempt() {
    const User = {
        name: "mister",
        password: "password",
    }
    const url = "localhost:8080/api/login";
    // const response = await fetch()

    let wrongPassword = document.getElementById("wrongPassword");
    let wrongPasswordStl = getComputedStyle(wrongPassword);
    if (wrongPasswordStl.display === "none") {
      wrongPassword.style.display = "block";
      let passwordForm = document.getElementById("passwordForm");
      passwordForm.style.marginBottom = 0;
    }
  }

console.log(x);

let x = 5;

