function fetchProducts() {
    fetch("/api/products") // Adjust the API endpoint accordingly
        .then((response) => response.json())
        .then((data) => {
            const productsContainer = document.getElementById("products");

            // Loop through products and create product cards
            data.forEach((product) => {
                const productCard = document.createElement("div");
                productCard.classList.add("product-card");
                productCard.innerHTML = `
                    <h2>${product.name}</h2>
                    <p>Price: $${product.price.toFixed(2)}</p>
                    <!-- Add more product details here -->
                `;
                productsContainer.appendChild(productCard);
            });
        })
        .catch((error) => {
            console.error("Error fetching products:", error);
        });
}

// Call the fetchProducts function when the page loads
window.addEventListener("load", fetchProducts);