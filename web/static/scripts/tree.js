/*
 * This file contains the functions that are used to toggle the visibility of the fields 
 * in the tree view of the tables.
 *
 * This file also contains the functions that are used to generate quick queries for the 
 * tables.
 */
function ToggleFields(id) {
  const fields = document.getElementById(`fields-${id}`);
  const button_svg = document.getElementById(`icon-${id}`);
  if (fields.classList.contains("hidden")) {
    fields.classList.remove("hidden");
    button_svg.setAttribute("transform", "rotate(0)");
  } else {
    fields.classList.add("hidden");
    button_svg.setAttribute("transform", "rotate(-90)");
  }
}


function LoadTableQuery(table) {
  const sql = document.getElementById("sql")
  sql.value = `SELECT * FROM ${table};`;
  sql.dispatchEvent(new Event("input", { bubbles: true }));
}
