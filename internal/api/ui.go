package api

const PortalHTML = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>Hospital Department Portal</title>
<style>
:root {
  color-scheme: light;
  --ink: #17202a;
  --muted: #5f6b76;
  --line: #d8dee4;
  --paper: #ffffff;
  --canvas: #f4f6f8;
  --primary: #006b5f;
  --primary-dark: #004f47;
  --danger: #b42318;
  --warning: #9a6700;
  --success: #16794b;
  --focus: #2f81f7;
}
* {
  box-sizing: border-box;
}
body {
  margin: 0;
  background: var(--canvas);
  color: var(--ink);
  font: 14px/1.45 system-ui, sans-serif;
}
button,
input,
select {
  font: inherit;
}
button {
  min-height: 36px;
  border: 1px solid var(--line);
  border-radius: 4px;
  background: var(--paper);
  color: var(--ink);
  cursor: pointer;
  padding: 6px 12px;
}
button:hover {
  border-color: var(--primary);
}
button:focus-visible,
input:focus-visible,
select:focus-visible {
  outline: 3px solid color-mix(in srgb, var(--focus) 25%, transparent);
  outline-offset: 1px;
}
button.primary {
  background: var(--primary);
  border-color: var(--primary);
  color: white;
}
button.primary:hover {
  background: var(--primary-dark);
}
header {
  align-items: center;
  background: var(--paper);
  border-bottom: 1px solid var(--line);
  display: flex;
  height: 58px;
  justify-content: space-between;
  padding: 0 24px;
}
.brand {
  align-items: center;
  display: flex;
  font-size: 17px;
  font-weight: 700;
  gap: 10px;
}
.brand-mark {
  align-items: center;
  background: var(--primary);
  border-radius: 4px;
  color: white;
  display: inline-flex;
  height: 30px;
  justify-content: center;
  width: 30px;
}
.shell {
  display: grid;
  grid-template-columns: 220px minmax(0, 1fr);
  min-height: calc(100vh - 58px);
}
nav {
  background: #20272d;
  color: white;
  padding: 18px 12px;
}
nav button {
  background: transparent;
  border-color: transparent;
  color: #dce3e8;
  display: block;
  margin-bottom: 4px;
  text-align: left;
  width: 100%;
}
nav button.active,
nav button:hover {
  background: #35414a;
  border-color: #46545f;
  color: white;
}
main {
  min-width: 0;
  padding: 24px;
}
.toolbar {
  align-items: center;
  display: flex;
  gap: 8px;
  justify-content: space-between;
  margin-bottom: 16px;
}
.toolbar h1 {
  font-size: 22px;
  margin: 0;
}
.filters {
  display: flex;
  gap: 8px;
}
input,
select {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 4px;
  color: var(--ink);
  min-height: 36px;
  padding: 6px 10px;
}
table {
  background: var(--paper);
  border: 1px solid var(--line);
  border-collapse: collapse;
  width: 100%;
}
th,
td {
  border-bottom: 1px solid var(--line);
  padding: 10px 12px;
  text-align: left;
  vertical-align: middle;
}
th {
  background: #f8fafb;
  color: var(--muted);
  font-size: 12px;
  font-weight: 650;
  text-transform: uppercase;
}
tbody tr:hover {
  background: #f7fbfa;
}
.status {
  align-items: center;
  display: inline-flex;
  font-size: 12px;
  font-weight: 650;
  gap: 6px;
}
.status::before {
  background: currentColor;
  border-radius: 50%;
  content: "";
  height: 7px;
  width: 7px;
}
.status.active,
.status.published {
  color: var(--success);
}
.status.pending,
.status.draft {
  color: var(--warning);
}
.status.suspended,
.status.cancelled,
.status.inactive {
  color: var(--danger);
}
.empty {
  color: var(--muted);
  padding: 54px 16px;
  text-align: center;
}
.panel {
  background: var(--paper);
  border: 1px solid var(--line);
  border-radius: 6px;
  margin-bottom: 16px;
}
.panel-heading {
  align-items: center;
  border-bottom: 1px solid var(--line);
  display: flex;
  justify-content: space-between;
  padding: 14px 16px;
}
.panel-heading h2 {
  font-size: 16px;
  margin: 0;
}
.panel-body {
  padding: 16px;
}
.details {
  display: grid;
  gap: 14px 24px;
  grid-template-columns: repeat(3, minmax(0, 1fr));
}
.detail-label {
  color: var(--muted);
  display: block;
  font-size: 12px;
  margin-bottom: 3px;
}
.detail-value {
  overflow-wrap: anywhere;
}
dialog {
  border: 1px solid var(--line);
  border-radius: 6px;
  box-shadow: 0 12px 40px rgb(0 0 0 / 18%);
  max-width: 520px;
  padding: 0;
  width: calc(100% - 32px);
}
dialog::backdrop {
  background: rgb(0 0 0 / 38%);
}
.dialog-heading {
  border-bottom: 1px solid var(--line);
  padding: 16px 20px;
}
.dialog-heading h2 {
  font-size: 17px;
  margin: 0;
}
form {
  padding: 18px 20px;
}
.form-grid {
  display: grid;
  gap: 14px;
  grid-template-columns: 1fr 1fr;
}
.field {
  display: grid;
  gap: 5px;
}
.field.full {
  grid-column: 1 / -1;
}
.field label {
  color: var(--muted);
  font-size: 12px;
  font-weight: 600;
}
.dialog-actions {
  border-top: 1px solid var(--line);
  display: flex;
  gap: 8px;
  justify-content: flex-end;
  margin: 18px -20px -18px;
  padding: 12px 20px;
}
.notice {
  border-left: 3px solid var(--focus);
  margin-bottom: 14px;
  padding: 10px 12px;
}
.notice.error {
  border-color: var(--danger);
  color: var(--danger);
}
.loading {
  color: var(--muted);
  padding: 30px;
  text-align: center;
}
.hidden {
  display: none !important;
}
@media (max-width: 860px) {
  .shell {
    grid-template-columns: 1fr;
  }
  nav {
    display: flex;
    gap: 4px;
    overflow-x: auto;
    padding: 8px;
  }
  nav button {
    margin: 0;
    white-space: nowrap;
    width: auto;
  }
  main {
    padding: 16px;
  }
  .details {
    grid-template-columns: 1fr 1fr;
  }
}
@media (max-width: 560px) {
  header {
    padding: 0 14px;
  }
  .toolbar {
    align-items: stretch;
    flex-direction: column;
  }
  .filters {
    display: grid;
    grid-template-columns: 1fr;
  }
  .form-grid,
  .details {
    grid-template-columns: 1fr;
  }
  table {
    display: block;
    overflow-x: auto;
  }
}
</style>
</head>
<body>
<header>
<div class="brand"><span class="brand-mark">H</span>Department Portal</div>
<div id="connection" class="status pending">Checking service</div>
</header>
<div class="shell">
<nav aria-label="Primary">
<button class="active" data-view="accounts">User management</button>
<button data-view="departments">Departments</button>
<button data-view="shifts">Duty schedule</button>
<button data-view="api">API reference</button>
</nav>
<main>
<section id="accounts-view">
<div class="toolbar">
<h1>User management</h1>
<div class="filters">
<select id="account-role" aria-label="Role filter">
<option value="">All roles</option>
<option value="doctor">Doctors</option>
<option value="nurse">Nurses</option>
<option value="administrator">Administrators</option>
</select>
<select id="account-status" aria-label="Status filter">
<option value="">All statuses</option>
<option value="pending">Pending</option>
<option value="active">Active</option>
<option value="suspended">Suspended</option>
</select>
<button class="primary" id="new-account">Add user</button>
</div>
</div>
<table>
<thead><tr><th>Name</th><th>Employee</th><th>Role</th><th>Department</th><th>Status</th><th>Contact</th></tr></thead>
<tbody id="account-rows"></tbody>
</table>
</section>
<section id="departments-view" class="hidden">
<div class="toolbar"><h1>Departments</h1><button class="primary" id="new-department">Add department</button></div>
<div id="department-list"></div>
</section>
<section id="shifts-view" class="hidden">
<div class="toolbar"><h1>Duty schedule</h1><button class="primary" id="new-shift">Add shift</button></div>
<table>
<thead><tr><th>Title</th><th>Department</th><th>Account</th><th>Starts</th><th>Ends</th><th>Status</th></tr></thead>
<tbody id="shift-rows"></tbody>
</table>
</section>
<section id="api-view" class="hidden">
<div class="panel">
<div class="panel-heading"><h2>API contract</h2><a href="/api/openapi.yaml">Open YAML</a></div>
<div class="panel-body">The OpenAPI 3.0 contract documents response envelopes, status mapping, and request identifiers.</div>
</div>
</section>
</main>
</div>
<dialog id="account-dialog">
<div class="dialog-heading"><h2>Add user account</h2></div>
<form id="account-form">
<div class="form-grid">
<div class="field"><label for="account-id">Account ID</label><input id="account-id" required></div>
<div class="field"><label for="employee-number">Employee number</label><input id="employee-number" required></div>
<div class="field full"><label for="display-name">Display name</label><input id="display-name" required></div>
<div class="field full"><label for="email">Email</label><input id="email" type="email" required></div>
<div class="field"><label for="role">Role</label><select id="role"><option value="doctor">Doctor</option><option value="nurse">Nurse</option><option value="administrator">Administrator</option></select></div>
<div class="field"><label for="department-id">Department ID</label><input id="department-id"></div>
</div>
<div class="dialog-actions"><button type="button" data-close>Cancel</button><button class="primary" type="submit">Create</button></div>
</form>
</dialog>
<script>
const state = {
  accounts: [],
  departments: [],
  shifts: [],
  view: "accounts"
};
const elements = {
  accountRows: document.querySelector("#account-rows"),
  departmentList: document.querySelector("#department-list"),
  shiftRows: document.querySelector("#shift-rows"),
  connection: document.querySelector("#connection")
};
function escapeHTML(value) {
  const node = document.createElement("span");
  node.textContent = String(value ?? "");
  return node.innerHTML;
}
function status(value) {
  return '<span class="status ' + escapeHTML(value) + '">' + escapeHTML(value) + '</span>';
}
async function request(path, options = {}) {
  const response = await fetch(path, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      "X-Request-ID": "portal-ui",
      ...(options.headers || {})
    }
  });
  const body = await response.json();
  if (!response.ok) {
    throw new Error(body.error?.message || "Request failed");
  }
  return body;
}
async function health() {
  try {
    await request("/healthz");
    elements.connection.className = "status active";
    elements.connection.textContent = "Service available";
  } catch (error) {
    elements.connection.className = "status suspended";
    elements.connection.textContent = "Service unavailable";
  }
}
async function loadAccounts() {
  const role = document.querySelector("#account-role").value;
  const statusValue = document.querySelector("#account-status").value;
  const query = new URLSearchParams();
  if (role) query.set("role", role);
  if (statusValue) query.set("status", statusValue);
  const body = await request("/api/accounts?" + query);
  state.accounts = body.items;
  renderAccounts();
}
function renderAccounts() {
  if (!state.accounts.length) {
    elements.accountRows.innerHTML = '<tr><td colspan="6" class="empty">No accounts match these filters.</td></tr>';
    return;
  }
  elements.accountRows.innerHTML = state.accounts.map(account => '<tr>' +
    '<td>' + escapeHTML(account.display_name) + '</td>' +
    '<td>' + escapeHTML(account.employee_number) + '</td>' +
    '<td>' + escapeHTML(account.role) + '</td>' +
    '<td>' + escapeHTML(account.department_id || "All departments") + '</td>' +
    '<td>' + status(account.status) + '</td>' +
    '<td>' + escapeHTML(account.email) + '</td>' +
    '</tr>').join("");
}
async function loadDepartments() {
  const body = await request("/api/departments");
  state.departments = body.items;
  renderDepartments();
}
function renderDepartments() {
  if (!state.departments.length) {
    elements.departmentList.innerHTML = '<div class="panel empty">No departments have been created.</div>';
    return;
  }
  elements.departmentList.innerHTML = state.departments.map(department => '<article class="panel">' +
    '<div class="panel-heading"><h2>' + escapeHTML(department.code + " · " + department.name) + '</h2>' + status(department.status) + '</div>' +
    '<div class="panel-body details">' +
    '<div><span class="detail-label">Location</span><span class="detail-value">' + escapeHTML(department.location || "Not set") + '</span></div>' +
    '<div><span class="detail-label">Phone</span><span class="detail-value">' + escapeHTML(department.phone || "Not set") + '</span></div>' +
    '<div><span class="detail-label">Email</span><span class="detail-value">' + escapeHTML(department.email || "Not set") + '</span></div>' +
    '<div><span class="detail-label">Head account</span><span class="detail-value">' + escapeHTML(department.head_account_id || "Not assigned") + '</span></div>' +
    '<div><span class="detail-label">Services</span><span class="detail-value">' + escapeHTML((department.services || []).join(", ") || "Not set") + '</span></div>' +
    '<div><span class="detail-label">Version</span><span class="detail-value">' + escapeHTML(department.version) + '</span></div>' +
    '</div></article>').join("");
}
async function loadShifts() {
  const body = await request("/api/shifts");
  state.shifts = body.items;
  renderShifts();
}
function renderShifts() {
  if (!state.shifts.length) {
    elements.shiftRows.innerHTML = '<tr><td colspan="6" class="empty">No duty shifts have been scheduled.</td></tr>';
    return;
  }
  elements.shiftRows.innerHTML = state.shifts.map(shift => '<tr>' +
    '<td>' + escapeHTML(shift.title) + '</td>' +
    '<td>' + escapeHTML(shift.department_id) + '</td>' +
    '<td>' + escapeHTML(shift.account_id) + '</td>' +
    '<td>' + escapeHTML(new Date(shift.start_at).toLocaleString()) + '</td>' +
    '<td>' + escapeHTML(new Date(shift.end_at).toLocaleString()) + '</td>' +
    '<td>' + status(shift.status) + '</td>' +
    '</tr>').join("");
}
async function showView(name) {
  state.view = name;
  document.querySelectorAll("main > section").forEach(section => section.classList.add("hidden"));
  document.querySelector("#" + name + "-view").classList.remove("hidden");
  document.querySelectorAll("nav button").forEach(button => button.classList.toggle("active", button.dataset.view === name));
  try {
    if (name === "accounts") await loadAccounts();
    if (name === "departments") await loadDepartments();
    if (name === "shifts") await loadShifts();
  } catch (error) {
    elements.connection.className = "status suspended";
    elements.connection.textContent = error.message;
  }
}
document.querySelectorAll("nav button").forEach(button => button.addEventListener("click", () => showView(button.dataset.view)));
document.querySelector("#account-role").addEventListener("change", loadAccounts);
document.querySelector("#account-status").addEventListener("change", loadAccounts);
document.querySelector("#new-account").addEventListener("click", () => document.querySelector("#account-dialog").showModal());
document.querySelectorAll("[data-close]").forEach(button => button.addEventListener("click", () => button.closest("dialog").close()));
document.querySelector("#account-form").addEventListener("submit", async event => {
  event.preventDefault();
  const payload = {
    id: document.querySelector("#account-id").value,
    display_name: document.querySelector("#display-name").value,
    employee_number: document.querySelector("#employee-number").value,
    email: document.querySelector("#email").value,
    role: document.querySelector("#role").value,
    department_id: document.querySelector("#department-id").value
  };
  try {
    await request("/api/accounts", {method: "POST", body: JSON.stringify(payload)});
    document.querySelector("#account-dialog").close();
    event.target.reset();
    await loadAccounts();
  } catch (error) {
    elements.connection.className = "status suspended";
    elements.connection.textContent = error.message;
  }
});
health();
showView("accounts");
</script>
</body>
</html>`
