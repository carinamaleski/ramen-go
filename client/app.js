const API_KEY = "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf";
const API_BASE_URL = "http://localhost:8080";
const BROTHS_ENDPOINT = `${API_BASE_URL}/broths`;
const PROTEINS_ENDPOINT = `${API_BASE_URL}/proteins`;
const ORDERS_ENDPOINT = `${API_BASE_URL}/orders`;

async function fetchData(endpoint) {
  console.log(`Fetching data from: ${endpoint}`);
  try {
    const response = await fetch(endpoint, {
      headers: {
        "x-api-key": API_KEY,
      },
    });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    console.log(
      `Data received from ${endpoint}:`,
      JSON.stringify(data, null, 2),
    );
    return data;
  } catch (error) {
    console.error(`Error fetching data from ${endpoint}:`, error);
    return null;
  }
}

async function getBroths() {
  try {
    const broths = await fetchData(BROTHS_ENDPOINT);
    const cardContainer = document.querySelector(
      ".content-cards .broth-card-container",
    );
    cardContainer.innerHTML = "";

    broths.forEach((broth) => {
      console.log("Broth data:", broth);
      const card = document.createElement("label");
      card.className = "card";
      card.innerHTML = `
        <input type="radio" name="broth" value="${broth.id}" class="card-radio">
        <div class="card-content">
          <img src="/public/icons/${broth.imageInactive}" alt="${broth.name}">
          <h3>${broth.name}</h3>
          <p class="desc">${broth.description}</p>
          <p class="price">US$ ${broth.price}</p>
        </div>
        `;
      cardContainer.appendChild(card);
    });
  } catch (error) {
    console.error("There was a problem fetching the broths:", error);
  }
}

async function getProteins() {
  try {
    const proteins = await fetchData(PROTEINS_ENDPOINT);
    const selectProtein = document.querySelector(
      ".content-cards .protein-card-container",
    );
    selectProtein.innerHTML = "";

    proteins.forEach((protein) => {
      console.log("Protein data:", protein);
      const card = document.createElement("label");
      card.className = "card";
      card.innerHTML = `
        <input type="radio" name="protein" value="${protein.id}" class="card-radio">
        <div class="card-content">
          <img src="/public/icons/${protein.imageInactive}" alt="${protein.name}">
          <h3>${protein.name}</h3>
          <p class="desc">${protein.description}</p>
          <p class="price">US$ ${protein.price}</p>
        </div>
      `;
      selectProtein.appendChild(card);
    });
  } catch (error) {
    console.error("There was a problem fetching the proteins:", error);
  }
}

async function createOrder() {
  const selectedBroth = document.querySelector('input[name="broth"]:checked');
  const selectedProtein = document.querySelector(
    'input[name="protein"]:checked',
  );

  if (!selectedBroth || !selectedProtein) {
    alert(
      "Please select both a broth and a protein before placing your order.",
    );
    return;
  }

  const order = {
    brothId: selectedBroth.value,
    proteinId: selectedProtein.value,
  };
  console.log("order:", order);
  try {
    const response = await fetch(ORDERS_ENDPOINT, {
      method: "POST",
      headers: {
        "x-api-key": API_KEY,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(order),
    });
    console.log("response:", response);
    if (response.ok) {
      const orderResult = await response.json();
      console.log("orderResult:", orderResult);
      alert(`Order placed successfully! Order ID: ${orderResult.id}`);
    } else {
      throw new Error("Failed to place order");
    }
  } catch (error) {
    console.error("Error placing order:", error);
    alert("Failed to place order. Please try again.");
  }
}

document.addEventListener("DOMContentLoaded", () => {
  getBroths();
  getProteins();
  document.addEventListener("change", function (event) {
    if (event.target.type === "radio") {
      const cards = document.querySelectorAll(".card");
      cards.forEach((card) => {
        const radio = card.querySelector('input[type="radio"]');
        const img = card.querySelector("img");
        if (radio.checked) {
          img.src = img.src.replace("imageInactive", "imageActive");
        } else {
          img.src = img.src.replace("imageActive", "imageInactive");
        }
      });
    }
  });
  const placeOrderBtn = document.getElementById("placeOrderBtn");
  placeOrderBtn.addEventListener("click", createOrder);
});
