import { useState } from "react";
import "./App.css";

function App() {
  const [name, setName] = useState("");
  const [phone, setPhone] = useState("");

  const [customers, setCustomers] = useState([]);

  function handleAddCustomer() {
    const newCustomer = {
      name: name,
      phone: phone,
    };

    setCustomers([
      ...customers, 
      newCustomer
    ]);

    // setName("");
    // setPhone("");
  }

  return (
    <>
      <input
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Name"
      />

      <input
        value={phone}
        onChange={(e) => setPhone(e.target.value)}
        placeholder="Phone"
      />

      <button onClick={handleAddCustomer}>
        Add Customer
      </button>

      <ul>
        {customers.map((customer, index) => (
          <li key={index}>
            {customer.name} - {customer.phone}
          </li>
        ))}
      </ul>
    </>
  );
}

export default App;
