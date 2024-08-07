/*
 * This file is used to toggle the password visibility in the password input field
 * of the new connection modal.
 */
const passwordInput = document.getElementById("db-password");
const toggleButton = document.getElementById("togglePassword");
const eyeIcon = document.getElementById("eyeIcon");
const urlInput = document.getElementById("db-url");

toggleButton.addEventListener("click", () => {
  if (passwordInput.type === "password") {
    passwordInput.type = "text";
    urlInput.type = "text";
    eyeIcon.setAttribute("stroke", "blue");
  } else {
    passwordInput.type = "password";
    urlInput.type = "password";
    eyeIcon.setAttribute("stroke", "currentColor");
  }
});
