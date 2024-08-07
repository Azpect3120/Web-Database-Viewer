/*
 * This file is used to generate the connection URL for the database connection
 * when the user changes inputs in the connection modal.
 *
 * This file is also used to show and hide the connection modal.
 */

// New connection modal
const modal = document.getElementById("connection-modal");

function ShowModal() {
  modal.classList.remove("hidden");
  modal.classList.remove("opacity-0");
}

function HideModal() {
  modal.classList.add("hidden");
  modal.classList.add("opacity-0");
}

const input = {
  host: document.getElementById("db-host"),
  port: document.getElementById("db-port"),
  driver: document.getElementById("db-driver"),
  username: document.getElementById("db-username"),
  password: document.getElementById("db-password"),
  database: document.getElementById("db-database"),
  connectionURL: document.getElementById("db-url")
}

// State for hiding the password in the connection URL
// let hidden = true;
// const togglePassword = () => {
//   hidden = !hidden;
//   GenerateURL(input);
// }

function GenerateURL(data) {
  // let password = hidden ? "●".repeat(data.password.value.length) : data.password.value;
  data.connectionURL.value = `${data.driver.value}://${data.username.value || "<username>"}:${data.password.value || "<password>"}@${data.host.value || "<host>"}:${data.port.value || "<port>"}/${data.database.value || "<database>"}?sslmode=disable`;
}

function ParseURL (data) {
  const regex = /^(?<protocol>[a-z]+):\/\/(?<username>[^:]+):(?<password>[^@]+)@(?<host>[^:]+):(?<port>[0-9]+)\/(?<database>[^\?]+)(\?(?<params>.*))?$/;
  const match = data.connectionURL.value.match(regex);

  if (match) {
    const { protocol, username, password, host, port, database } = match.groups;

    data.host.value = host;
    data.port.value = port;
    data.password.value = password;
    data.username.value = username;
    data.driver.value = protocol;
    data.database.value = database;
    
    data.connectionURL.classList.remove("text-red-500");
    document.getElementById("db-url-invalid").classList.add("hidden");

  } else {
    data.connectionURL.classList.add("text-red-500");
    document.getElementById("db-url-invalid").classList.remove("hidden");
  }
}

// Create the event listeners for the input fields
for (const key in input) {
  if (key == "connectionURL") {
    input[key].addEventListener("input", () => {
      ParseURL(input);
    })
  } else {
    input[key].addEventListener("input", () => {
      GenerateURL(input);
    })
  }
}
