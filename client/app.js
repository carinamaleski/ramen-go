const API_BASE_URL = "https://localhost:8080";
const BROTHS_ENDPOINT = `${API_BASE_URL}/broths`;
const PROTEINS_ENDPOINT = `${API_BASE_URL}/proteins`;
const ORDERS_ENDPOINT = `${API_BASE_URL}/orders`;

async function getBroths() {
  const response = await fetch(BROTHS_ENDPOINT);
  const broths = await response.json();
  const selectBroth = document.getElementById("broth");

  broths.forEach((broth) => {
    const option = document.createElement("option");
    option.value = broth.id;
    option.textContent = broth.name;
    option.textContent = broth.description;
    selectBroth.appendChild(option);
  });
}

async function getProteins() {
  const response = await fetch(PROTEINS_ENDPOINT);
  const proteins = await response.json();
  const selectProtein = document.getElementById("protein");

  proteins.forEach((protein) => {
    const option = document.createElement("option");
    option.value = protein.id;
    option.textContent = protein.name;
    option.textContent = protein.description;
    selectProtein.appendChild(option);
  });
}

async function createOrder() {
  const brothId = document.getElementById("broth").value;
  const proteinId = document.getElementById("protein").value;

  const order = {
    brothId: brothId,
    proteinId: proteinId,
  };

  const response = await fetch(ORDERS_ENDPOINT, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(order),
  });

  if (response.ok) {
    const order = await response.json();
    document.getElementById("order").textContent =
      `Pedido criado com sucesso! ID do pedido: ${order.id}`;
  } else {
    document.getElementById("order").textContent =
      "Erro ao criar o pedido. Tente novamente.";
  }
}

document.addEventListener("DOMContentLoaded", () => {
  getBroths();
  getProteins();

  document.getElementById("submit").addEventListener("click", createOrder);
});
